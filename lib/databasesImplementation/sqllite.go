package databasesImplementation

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/tobedeterminedhq/tbd/lib/databasesImplementation/databaseImplementationBase"

	_ "github.com/glebarez/go-sqlite"
)

type SQLLite struct {
	DB *sql.DB
}

func NewSqlLite(filePath string) (*SQLLite, error) {
	db, err := sql.Open("sqlite", filePath+"?_pragma=foreign_keys(1)")
	if err != nil {
		return nil, err
	}
	return &SQLLite{DB: db}, nil
}

// NewSqlLiteInMemory returns an in-memory sqlite instance through the call of NewSqlLite
func NewSqlLiteInMemory() (*SQLLite, error) {
	return NewSqlLite(":memory:")
}

// ListTables implements the interface Database
func (db *SQLLite) ListTables(ctx context.Context) (tables []string, err error) {
	rows, err := db.DB.QueryContext(ctx, "SELECT name FROM sqlite_master WHERE type='table'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		tables = append(tables, name)
	}
	return tables, nil
}

// ListViews implements the interface Database
func (db *SQLLite) ListViews(ctx context.Context) (views []string, err error) {
	rows, err := db.DB.QueryContext(ctx, "SELECT name FROM sqlite_master WHERE type='view'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		views = append(views, name)
	}
	return views, nil
}

func (db *SQLLite) ListColumns(ctx context.Context, table string) (columns []string, err error) {
	doesExist, err := db.doesTableExist(ctx, table)
	if err != nil {
		return nil, err
	}
	if !doesExist {
		return nil, fmt.Errorf("table %s does not exist", table)
	}
	rows, err := db.DB.QueryContext(ctx, "PRAGMA table_info("+table+")")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var d1, d2, d3, d4, d5 interface{}
		if err := rows.Scan(&d1, &name, &d2, &d3, &d4, &d5); err != nil {
			return nil, err
		}
		columns = append(columns, name)
	}
	return columns, nil
}

func (db *SQLLite) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.DB.ExecContext(ctx, query, args...)
}

func (db *SQLLite) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return db.DB.QueryContext(ctx, query, args...)
}

func (db *SQLLite) doesTableExist(ctx context.Context, table string) (bool, error) {
	rows, err := db.DB.QueryContext(ctx, "SELECT name FROM sqlite_master WHERE type='table' AND name=?", table)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	return rows.Next(), nil
}

func (db *SQLLite) Close(ctx context.Context) error {
	return db.DB.Close()
}

func (db *SQLLite) ReturnFullPathRequirement(tableName string) string {
	return tableName
}

func (db *SQLLite) SeedsDropTableQuery(tableName string) string {
	return databaseImplementationBase.BaseForSeedsDeleteTable(tableName)
}

func (db *SQLLite) SeedsCreateTableQuery(tableName string, columns []string) (string, error) {
	return databaseImplementationBase.BaseForSeedsCreateTable(tableName, columns)
}

func (db *SQLLite) SeedsInsertIntoTableQuery(tableName string, columns []string, values [][]string) (string, error) {
	return databaseImplementationBase.BaseForSeedsInsertTable(tableName, columns, values)
}
