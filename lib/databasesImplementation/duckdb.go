package databasesImplementation

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/benfdking/tbd/go/lib/databasesImplementation/databaseImplementationBase"

	_ "github.com/marcboeker/go-duckdb"
)

type DuckDB struct {
	DB *sql.DB
}

func (db *DuckDB) ReturnFullPathRequirement(tableName string) string {
	// TODO implement me
	panic("implement me")
}

func (db *DuckDB) SeedsDropTableQuery(tableName string) string {
	return databaseImplementationBase.BaseForSeedsDeleteTable(tableName)
}

func (db *DuckDB) SeedsCreateTableQuery(tableName string, columns []string) (string, error) {
	return databaseImplementationBase.BaseForSeedsCreateTable(tableName, columns)
}

func (db *DuckDB) SeedsInsertIntoTableQuery(tableName string, columns []string, values [][]string) (string, error) {
	return databaseImplementationBase.BaseForSeedsInsertTable(tableName, columns, values)
}

func NewDuckDB(filePath string, params map[string]string) (*DuckDB, error) {
	db, err := sql.Open("duckdb", filePath+"?"+mapToParams(params))
	if err != nil {
		return nil, err
	}
	return &DuckDB{DB: db}, nil
}

func mapToParams(params map[string]string) string {
	if len(params) == 0 {
		return ""
	}
	out := ""
	for k, v := range params {
		out += k + "=" + v + "&"
	}
	return out[:len(out)-1]
}

// ListTables implements the interface Database
func (db *DuckDB) ListTables(ctx context.Context) (tables []string, err error) {
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
func (db *DuckDB) ListViews(ctx context.Context) (views []string, err error) {
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

func (db *DuckDB) ListColumns(ctx context.Context, table string) (columns []string, err error) {
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

func (db *DuckDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.DB.ExecContext(ctx, query, args...)
}

func (db *DuckDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return db.DB.QueryContext(ctx, query, args...)
}

func (db *DuckDB) doesTableExist(ctx context.Context, table string) (bool, error) {
	rows, err := db.DB.QueryContext(ctx, "SELECT name FROM sqlite_master WHERE type='table' AND name=?", table)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	return rows.Next(), nil
}

func (db *DuckDB) Close(ctx context.Context) error {
	return db.DB.Close()
}
