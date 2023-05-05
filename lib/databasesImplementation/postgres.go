package databasesImplementation

import (
	"context"
	"database/sql"

	"github.com/tobedeterminedhq/tbd/lib/databasesImplementation/databaseImplementationBase"

	_ "github.com/jackc/pgx/v5"
	"github.com/tobedeterminedhq/tbd/lib/databases"
)

func (p Postgres) SeedsDropTableQuery(tableName string) string {
	return databaseImplementationBase.BaseForSeedsDeleteTable(tableName)
}

func (p Postgres) SeedsCreateTableQuery(tableName string, columns []string) (string, error) {
	return databaseImplementationBase.BaseForSeedsCreateTable(tableName, columns)
}

func (p Postgres) SeedsInsertIntoTableQuery(tableName string, columns []string, values [][]string) (string, error) {
	return databaseImplementationBase.BaseForSeedsInsertTable(tableName, columns, values)
}

func NewPostgres(connectionString string) (databases.Database, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return Postgres{db: db}, nil
}

type Postgres struct {
	db *sql.DB
}

func (p Postgres) ListTables(ctx context.Context) ([]string, error) {
	const query = "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND schemaname != 'information_schema';"

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var views []string
	for rows.Next() {
		var view string
		if err := rows.Scan(&view); err != nil {
			return nil, err
		}
		views = append(views, view)
	}

	return views, nil
}

func (p Postgres) ListViews(ctx context.Context) ([]string, error) {
	const query = "SELECT table_name FROM INFORMATION_SCHEMA.views WHERE table_schema = ANY (current_schemas(false))"

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var views []string
	for rows.Next() {
		var view string
		if err := rows.Scan(&view); err != nil {
			return nil, err
		}
		views = append(views, view)
	}

	return views, nil
}

func (p Postgres) ListColumns(ctx context.Context, table string) ([]string, error) {
	// TODO Will need to check if the table exists
	// TODO Will need to
	const query = "SELECT column_name FROM INFORMATION_SCHEMA.COLUMNS WHERE table_name = $1"

	rows, err := p.db.QueryContext(ctx, query, table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []string
	for rows.Next() {
		var column string
		if err := rows.Scan(&column); err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}

	return columns, nil
}

func (p Postgres) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return p.db.ExecContext(ctx, query, args...)
}

func (p Postgres) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return p.db.QueryContext(ctx, query, args...)
}

func (p Postgres) Close(ctx context.Context) error {
	return p.db.Close()
}

func (p Postgres) ReturnFullPathRequirement(tableName string) string {
	return tableName
}
