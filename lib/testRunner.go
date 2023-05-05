package lib

import (
	"context"
	"fmt"

	"github.com/tobedeterminedhq/tbd/lib/databases"
)

// RunTestSql runs a test sql statement and returns an error if the test fails
func RunTestSql(ctx context.Context, db databases.Database, sql string) error {
	rows, err := db.QueryContext(ctx, sql)
	if err != nil {
		return fmt.Errorf("error running test with sql '%s': %w", sql, err)
	}
	defer rows.Close()

	var outs [][]string
	for rows.Next() {
		columns, err := rows.Columns()
		if err != nil {
			return fmt.Errorf("error getting columns with sql '%s': %w", sql, err)
		}

		out := make([]string, len(columns))
		outPtrs := make([]interface{}, len(columns))

		for i := range columns {
			outPtrs[i] = &out[i]
		}
		err = rows.Scan(outPtrs...)
		if err != nil {
			return fmt.Errorf("error scanning result with sql '%s': %w", sql, err)
		}

		outs = append(outs, out)
	}
	if len(outs) > 0 {
		// TODO Find a way to print the outs
		return fmt.Errorf("error with sql '%s', expected 0 results, got %d: %w", sql, 0, err)
	}
	return nil
}
