package lib

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/benfdking/tbd/go/lib/databases"
	"github.com/samber/lo"
)

// ParseTableSchemaSeeds returns the SQL statements to create a table and insert the data from a CSV file.
// doNotIncludeData is used to generate only the create table statement without any inserts.
func ParseTableSchemaSeeds(database databases.Database, tableName string, reader io.Reader, doNotIncludeData bool) ([]string, error) {
	r := csv.NewReader(reader)
	all, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("read file to table '%s': %w", tableName, err)
	}
	groups := lo.GroupBy[[]string, int](all, func(s []string) int { return len(s) })
	if len(groups) != 1 {
		return nil, fmt.Errorf("expect all rows to have same length")
	}
	dropStatement := database.SeedsDropTableQuery(tableName)
	createStatement, err := database.SeedsCreateTableQuery(tableName, all[0])
	if err != nil {
		return nil, err
	}
	if len(all) == 1 || doNotIncludeData {
		return []string{dropStatement, createStatement}, nil
	}
	insertStatement, err := database.SeedsInsertIntoTableQuery(tableName, all[0], all[1:])
	if err != nil {
		return nil, err
	}
	return []string{dropStatement, createStatement, insertStatement}, nil
}
