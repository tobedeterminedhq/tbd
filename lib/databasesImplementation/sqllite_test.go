package databasesImplementation_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tobedeterminedhq/tbd/lib/databasesImplementation"
)

// TODO Should make a test suite for this that is tested on every database

func TestSQLLite_ListColumns(t *testing.T) {
	tests := []struct {
		name        string
		prepFunc    func(t *testing.T, db *databasesImplementation.SQLLite) error
		table       string
		wantColumns []string
		wantErr     assert.ErrorAssertionFunc
	}{
		{
			name: "error table does not exist",
			prepFunc: func(t *testing.T, db *databasesImplementation.SQLLite) error {
				return nil
			},
			table:       "does_not_exist",
			wantColumns: nil,
			wantErr:     assert.Error,
		},
		{
			name: "table does exist",
			prepFunc: func(t *testing.T, db *databasesImplementation.SQLLite) error {
				_, err := db.ExecContext(context.Background(), "CREATE TABLE does_exist (id INTEGER PRIMARY KEY, name TEXT)")
				return err
			},
			table:       "does_exist",
			wantColumns: []string{"id", "name"},
			wantErr:     assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			db, err := databasesImplementation.NewSqlLiteInMemory()
			require.NoError(t, err)

			err = tt.prepFunc(t, db)
			require.NoError(t, err)

			gotColumns, err := db.ListColumns(ctx, tt.table)

			tt.wantErr(t, err)
			assert.Equal(t, tt.wantColumns, gotColumns)
		})
	}
}
