package databases

import (
	"context"
	"database/sql"
)

type Database interface {
	ListTables(ctx context.Context) ([]string, error)
	ListViews(ctx context.Context) ([]string, error)
	// ListColumns returns the columns of a table in the order they are defined in the table.
	// If the table does not exist, an error is returned.
	ListColumns(ctx context.Context, table string) ([]string, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	// Close the connection to the database. Easily used with defer.
	// TODO Make sure this is called in the main functions.
	Close(ctx context.Context) error
	// For Seeds section

	// SeedsDropTableQuery drops a table if it exists.
	SeedsDropTableQuery(tableName string) string
	// SeedsCreateTableQuery drops a table if it exists where the columns are Text/String equivalent.
	SeedsCreateTableQuery(tableName string, columns []string) (string, error)
	// SeedsInsertIntoTableQuery inserts values into a table.
	SeedsInsertIntoTableQuery(tableName string, columns []string, values [][]string) (string, error)
	// For Models section

	// ReturnFullPathRequirement takes in the name of the target table and prefixes it with any necessary schema/paths
	// to make it a full path.
	ReturnFullPathRequirement(tableName string) string
}
