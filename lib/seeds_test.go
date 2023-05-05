package lib_test

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/benfdking/tbd/go/lib"
	"github.com/benfdking/tbd/go/lib/databasesImplementation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultSql_ParseTableSchemaSeeds(t *testing.T) {
	tests := []struct {
		name               string
		tableName          string
		doNotIncludeInsert bool
		reader             io.Reader
		want               []string
	}{
		{
			name:               "simple example with 2 rows",
			tableName:          "sample_table",
			reader:             strings.NewReader("column_1,column_2\nvalue_1,value_2\nvalue_3,value_4"),
			doNotIncludeInsert: false,
			want: []string{
				`DROP TABLE IF EXISTS sample_table`,
				`CREATE TABLE sample_table (column_1 TEXT,column_2 TEXT)`,
				`INSERT INTO sample_table (column_1,column_2) VALUES ('value_1','value_2'),('value_3','value_4')`,
			},
		},
		{
			name:               "simple example with 2 rows, skip insert",
			tableName:          "sample_table",
			reader:             strings.NewReader("column_1,column_2\nvalue_1,value_2\nvalue_3,value_4"),
			doNotIncludeInsert: true,
			want: []string{
				`DROP TABLE IF EXISTS sample_table`,
				`CREATE TABLE sample_table (column_1 TEXT,column_2 TEXT)`,
			},
		},
		{
			name:               "simple example with 0 rows",
			tableName:          "sample_table",
			reader:             strings.NewReader("column_1,column_2\n"),
			doNotIncludeInsert: false,
			want: []string{
				`DROP TABLE IF EXISTS sample_table`,
				`CREATE TABLE sample_table (column_1 TEXT,column_2 TEXT)`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := databasesImplementation.NewSqlLiteInMemory()
			require.NoError(t, err)

			got, err := lib.ParseTableSchemaSeeds(db, tt.tableName, tt.reader, tt.doNotIncludeInsert)

			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestDefaultSql_ParseTableSchemaSeeds_Reapply tests the reapplication of seeds to ensure they change the database
// correctly consistently. The test applies various types of "reapplications".
func TestDefaultSql_ParseTableSchemaSeeds_Reapply(t *testing.T) {
	t.Parallel()

	const original = `number,name
1,tom
2,jerry
3,david`
	const tableName = "test_table"

	tests := []struct {
		name         string
		replacingCSV string
		checkSQL     string
		wantValues   [][]string
	}{
		{
			name:         "same thing",
			replacingCSV: original,
			checkSQL:     "SELECT number AS TEXT, name FROM test_table ORDER BY number",
			wantValues: [][]string{
				{"1", "tom"},
				{"2", "jerry"},
				{"3", "david"},
			},
		},
		{
			name:         "added row",
			replacingCSV: original + "\n4,peter",
			checkSQL:     "SELECT number AS TEXT, name FROM test_table ORDER BY number",
			wantValues: [][]string{
				{"1", "tom"},
				{"2", "jerry"},
				{"3", "david"},
				{"4", "peter"},
			},
		},
		{
			name: "added column",
			replacingCSV: `number,name,last_name
1,tom,peters
2,jerry,smith
3,david,seagull`,
			checkSQL: "SELECT number AS TEXT, name, last_name FROM test_table ORDER BY number",
			wantValues: [][]string{
				{"1", "tom", "peters"},
				{"2", "jerry", "smith"},
				{"3", "david", "seagull"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			// Database setup
			db, err := databasesImplementation.NewSqlLiteInMemory()
			require.NoError(t, err)

			// Apply original
			{
				sql, err := lib.ParseTableSchemaSeeds(db, tableName, strings.NewReader(original), false)
				require.NoError(t, err)

				_, err = db.ExecContext(ctx, strings.Join(sql, ";"))
				require.NoError(t, err)
			}

			// Check original value
			{
				rows, err := db.QueryContext(ctx, "SELECT number AS TEXT, name FROM test_table ORDER BY number")
				require.NoError(t, err)
				var values [][2]string
				for rows.Next() {
					var value [2]string
					err := rows.Scan(&value[0], &value[1])
					require.NoError(t, err)
					values = append(values, value)
				}

				require.Equal(t,
					[][2]string{
						{"1", "tom"},
						{"2", "jerry"},
						{"3", "david"},
					}, values)
			}

			// Apply new
			{
				sql, err := lib.ParseTableSchemaSeeds(db, tableName, strings.NewReader(tt.replacingCSV), false)
				require.NoError(t, err)

				_, err = db.ExecContext(ctx, strings.Join(sql, ";"))
				require.NoError(t, err)
			}

			// Check new values
			{
				rows, err := db.QueryContext(ctx, tt.checkSQL)
				require.NoError(t, err)

				var outs [][]string
				for rows.Next() {
					var out []string

					columns, err := rows.Columns()
					require.NoError(t, err)

					out = make([]string, len(columns))
					outPtrs := make([]interface{}, len(columns))
					for i := range columns {
						outPtrs[i] = &out[i]
					}
					err = rows.Scan(outPtrs...)
					require.NoError(t, err)

					outs = append(outs, out)
				}

				assert.Equal(t, tt.wantValues, outs)
			}
		})
	}
}
