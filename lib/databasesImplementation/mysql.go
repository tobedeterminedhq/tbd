package databasesImplementation

import (
	"context"
	"database/sql"

	"github.com/benfdking/tbd/go/lib/databasesImplementation/databaseImplementationBase"

	"github.com/benfdking/tbd/go/lib/databases"
	_ "github.com/go-sql-driver/mysql"
)

func (m Mysql) SeedsDropTableQuery(tableName string) string {
	return databaseImplementationBase.BaseForSeedsDeleteTable(tableName)
}

func (m Mysql) SeedsCreateTableQuery(tableName string, columns []string) (string, error) {
	return databaseImplementationBase.BaseForSeedsCreateTable(tableName, columns)
}

func (m Mysql) SeedsInsertIntoTableQuery(tableName string, columns []string, values [][]string) (string, error) {
	return databaseImplementationBase.BaseForSeedsInsertTable(tableName, columns, values)
}

func NewMySql(connectionString string) (databases.Database, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return Mysql{db: db}, nil
}

type Mysql struct {
	db *sql.DB
}

func (m Mysql) ReturnFullPathRequirement(tableName string) string {
	return tableName
}

func (m Mysql) ListTables(ctx context.Context) ([]string, error) {
	rows, err := m.db.QueryContext(ctx, "SHOW TABLES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tables := make([]string, 0)
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, nil
}

func (m Mysql) ListViews(ctx context.Context) ([]string, error) {
	rows, err := m.db.QueryContext(ctx, "SHOW VIEWS")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tables := make([]string, 0)
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, nil
}

func (m Mysql) ListColumns(ctx context.Context, table string) ([]string, error) {
	// TODO Should do a proper substitution here
	rows, err := m.db.QueryContext(ctx, "SHOW COLUMNS FROM "+table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns := make([]string, 0)
	for rows.Next() {
		var column string
		if err := rows.Scan(&column); err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}
	return columns, nil
}

func (m Mysql) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return m.db.ExecContext(ctx, query, args...)
}

func (m Mysql) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return m.db.QueryContext(ctx, query, args...)
}

func (m Mysql) Close(ctx context.Context) error {
	return m.db.Close()
}
