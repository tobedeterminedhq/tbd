package lib_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tobedeterminedhq/tbd/lib"
	"github.com/tobedeterminedhq/tbd/lib/databasesImplementation"
	servicev1 "github.com/tobedeterminedhq/tbd/proto_gen/go/tbd/service/v1"
)

func TestParseModelSchemasToViews(t *testing.T) {
	t.Parallel()

	type args struct {
		fileReader            io.Reader
		tableName             string
		nameReplacingStrategy func(name string) string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "simple example",
			args: args{
				fileReader:            bytes.NewBufferString("SELECT * FROM tbd.users"),
				tableName:             "view_name",
				nameReplacingStrategy: lib.ReplaceReferenceStringFound("tbd", map[string]*servicev1.Source{}),
			},
			want:    "DROP VIEW IF EXISTS view_name; CREATE VIEW view_name AS SELECT * FROM users;",
			wantErr: assert.NoError,
		},
		{
			name: "simple example that also has a source",
			args: args{
				fileReader: bytes.NewBufferString("SELECT * FROM tbd.users"),
				tableName:  "view_name",
				nameReplacingStrategy: lib.ReplaceReferenceStringFound("tbd", map[string]*servicev1.Source{
					"users": {
						Name:     "users",
						Path:     "schema.users_123",
						FilePath: "models/test.yaml",
						Columns:  nil,
					},
				}),
			},
			want:    "DROP VIEW IF EXISTS view_name; CREATE VIEW view_name AS SELECT * FROM schema.users_123;",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lib.ParseModelSchemasToViews(tt.args.fileReader, tt.args.tableName, "tbd", tt.args.nameReplacingStrategy)
			if !tt.wantErr(t, err, fmt.Sprintf("ParseModelSchemasToViews(%v, %v)", tt.args.fileReader, tt.args.tableName)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ParseModelSchemasToViews(%v, %v, %v)", tt.args.fileReader, tt.args.tableName, tt.args.nameReplacingStrategy)
		})
	}
}

// TestParseModelSchemasToViews_ReapplyingModel tests that the model can be reapplied to the database
// without error and the updated data be returned correctly to an in-memory SQLLite database.
func TestParseModelSchemasToViews_ReapplyingModelSqlLite(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	// Setting up database
	db, err := databasesImplementation.NewSqlLiteInMemory()
	require.NoError(t, err)

	// Setting up base table to query from view
	_, err = db.ExecContext(ctx, "CREATE TABLE users (id INTEGER PRIMARY KEY, user_name TEXT);")

	// Applying the view first time around
	const viewName = "reused_view_name"
	const firstQuery = "SELECT id AS id_int, user_name as name FROM users"

	model, err := lib.ParseModelSchemasToViews(
		bytes.NewBufferString(firstQuery),
		viewName,
		"tbd",
		func(name string) string { return name },
	)
	require.NoError(t, err)
	_, err = db.ExecContext(ctx, model)
	require.NoError(t, err)
	_, err = db.ExecContext(ctx, "SELECT id_int, name FROM reused_view_name")
	require.NoError(t, err)

	// Applying the same view and checking it
	model, err = lib.ParseModelSchemasToViews(
		bytes.NewBufferString(firstQuery),
		viewName,
		"tbd",
		func(name string) string { return name },
	)
	require.NoError(t, err)
	_, err = db.ExecContext(ctx, model)
	require.NoError(t, err)
	_, err = db.ExecContext(ctx, "SELECT id_int, name FROM reused_view_name")
	require.NoError(t, err)

	// Slightly checking view and checking it
	model, err = lib.ParseModelSchemasToViews(
		bytes.NewBufferString("SELECT id AS id_int_again, user_name as name FROM users"),
		viewName,
		"tbd",
		func(name string) string { return name },
	)
	require.NoError(t, err)
	_, err = db.ExecContext(ctx, model)
	require.NoError(t, err)
	_, err = db.ExecContext(ctx, "SELECT id_int, name FROM reused_view_name")
	require.Error(t, err)
	_, err = db.ExecContext(ctx, "SELECT id_int_again, name FROM reused_view_name")
	require.NoError(t, err)
}

func BenchmarkParseModelSchemasToViews(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := lib.ParseModelSchemasToViews(
			bytes.NewBufferString("SELECT * FROM tbd.users"),
			"view_name",
			"tbd",
			func(name string) string { return name },
		)
		if err != nil {
			b.Fatal(err)
		}
	}
}
