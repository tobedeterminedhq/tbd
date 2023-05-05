package lib

import (
	"context"
	"testing"

	"github.com/benfdking/tbd/go/lib/databases"
	"github.com/benfdking/tbd/go/lib/databasesImplementation"
	"github.com/stretchr/testify/require"

	servicev1 "github.com/benfdking/tbd/proto/gen/go/tbd/service/v1"

	"github.com/stretchr/testify/assert"
)

func TestGenerateNotNullSqlColumnTest(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		test *servicev1.TestNotNull
		want string
	}{
		{
			name: "simple example",
			test: &servicev1.TestNotNull{
				Model:  "users",
				Path:   "users_123",
				Column: "id",
			},
			want: "SELECT * FROM users_123 WHERE id IS NULL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateTestSqlNotNull(tt.test)
			assert.Equalf(t, tt.want, got, "GenerateTestSqlNotNull(%v)", tt.test)
		})
	}
}

func BenchmarkGenerateNotNullSqlColumnTest(b *testing.B) {
	test := &servicev1.TestNotNull{
		Model:  "users",
		Column: "id",
	}

	for n := 0; n < b.N; n++ {
		GenerateTestSqlNotNull(test)
	}
}

func TestGenerateUniqueSQLColumnTest(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		test *servicev1.TestUnique
		want string
	}{
		{
			name: "simple example",
			test: &servicev1.TestUnique{
				Model:  "users",
				Path:   "users_123",
				Column: "id",
			},
			want: `SELECT * FROM (
    SELECT id
    FROM users_123
    WHERE id IS NOT NULL
    GROUP BY id
    HAVING count(*) > 1
)`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateTestSqlUnique(tt.test)
			assert.Equalf(t, tt.want, got, "GenerateTestSqlNotNull(%v)", tt.test)
		})
	}
}

func BenchmarkGenerateTestSqlUnique(b *testing.B) {
	test := &servicev1.TestUnique{
		Model:  "users",
		Column: "id",
	}

	for n := 0; n < b.N; n++ {
		GenerateTestSqlUnique(test)
	}
}

func TestGenerateTestNameRelationship(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		test *servicev1.TestRelationship
		want string
	}{
		{
			name: "simple example",
			test: &servicev1.TestRelationship{
				FilePath:     "test/path.yaml",
				SourceModel:  "users",
				SourceColumn: "id",
				TargetModel:  "user_duplicated",
				TargetColumn: "user_dup_id",
			},
			want: "test_users_id_relationship_user_duplicated_user_dup_id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GenerateTestNameRelationship(tt.test), "GenerateTestNameRelationship(%v)", tt.test)
		})
	}
}

func TestGenerateTestSqlRelationship(t *testing.T) {
	tests := []struct {
		name string
		test *servicev1.TestRelationship
		want string
	}{
		{
			name: "simple example",
			test: &servicev1.TestRelationship{
				FilePath:     "test/path.yaml",
				SourceModel:  "users",
				SourcePath:   "users_123",
				SourceColumn: "id",
				TargetModel:  "user_duplicated",
				TargetPath:   "user_duplicated_123",
				TargetColumn: "user_dup_id",
			},
			want: "SELECT * FROM users_123 WHERE id IS NOT NULL AND id NOT IN (SELECT user_dup_id FROM user_duplicated_123)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateTestSqlRelationship(tt.test)
			assert.Equalf(t, tt.want, got, "GenerateTestSqlRelationship(%v)", tt.test)
		})
	}
}

func TestGenerateTestNameAcceptedValues(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		test *servicev1.TestAcceptedValues
		want string
	}{
		// TODO: Add test cases.
		{
			name: "simple example",
			test: &servicev1.TestAcceptedValues{
				FilePath:       "test/path.yaml",
				Model:          "users",
				Column:         "id",
				AcceptedValues: []string{"1", "2", "3"},
			},
			want: "test_users_id_accepted_values",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GenerateTestNameAcceptedValues(tt.test), "GenerateTestNameAcceptedValues(%v)", tt.test)
		})
	}
}

func TestGenerateTestSqlAcceptedValues(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		test *servicev1.TestAcceptedValues
		want string
	}{
		{
			name: "simple example",
			test: &servicev1.TestAcceptedValues{
				FilePath:       "test/path.yaml",
				Model:          "users",
				Path:           "users",
				Column:         "id",
				AcceptedValues: []string{"1", "2", "3"},
			},
			want: "SELECT * FROM users WHERE id IS NOT NULL AND id NOT IN ('1','2','3')",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateTestSqlAcceptedValues(tt.test)

			assert.Equalf(t, tt.want, got, "GenerateTestSqlAcceptedValues(%v)", tt.test)
		})
	}
}

// TestGenerateTestSqlAcceptedValues_ActuallyWorks tests that the generated SQL actually works.
// TODO This shall eventually need to be changed to work with all the databases.
func TestGenerateTestSqlAcceptedValues_ActuallyWorks(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		databasePrep func(context context.Context, db databases.Database) error
		test         *servicev1.TestAcceptedValues
		wantRows     int
	}{
		{
			name: "simple example, valid values",
			databasePrep: func(ctx context.Context, db databases.Database) error {
				_, err := db.ExecContext(ctx, "CREATE TABLE users_123 (id INT)")
				if err != nil {
					return err
				}
				_, err = db.ExecContext(ctx, "INSERT INTO users_123 (id) VALUES (1), (2), (3)")
				if err != nil {
					return err
				}
				return nil
			},
			test: &servicev1.TestAcceptedValues{
				FilePath:       "test/path.yaml",
				Model:          "users",
				Path:           "users_123",
				Column:         "id",
				AcceptedValues: []string{"1", "2", "3"},
			},
			wantRows: 0,
		},
		{
			name: "2 invalid values",
			databasePrep: func(ctx context.Context, db databases.Database) error {
				_, err := db.ExecContext(ctx, "CREATE TABLE users_123 (id INT)")
				if err != nil {
					return err
				}
				_, err = db.ExecContext(ctx, "INSERT INTO users_123 (id) VALUES (1), (2), (3), (4), (5)")
				if err != nil {
					return err
				}
				return nil
			},
			test: &servicev1.TestAcceptedValues{
				FilePath:       "test/path.yaml",
				Model:          "users",
				Path:           "users_123",
				Column:         "id",
				AcceptedValues: []string{"1", "2", "3"},
			},
			wantRows: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			db, err := databasesImplementation.NewSqlLiteInMemory()
			require.NoError(t, err)

			err = tt.databasePrep(ctx, db)
			require.NoError(t, err)

			got := GenerateTestSqlAcceptedValues(tt.test)

			rows, err := db.QueryContext(ctx, got)
			require.NoError(t, err)

			defer rows.Close()
			count := 0
			for rows.Next() {
				count += 1
			}

			assert.Equal(t, tt.wantRows, count)
		})
	}
}
