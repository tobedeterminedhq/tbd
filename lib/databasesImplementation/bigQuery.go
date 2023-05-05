package databasesImplementation

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/tobedeterminedhq/tbd/lib/databasesImplementation/databaseImplementationBase"

	_ "gorm.io/driver/bigquery/driver"
)

// TODO Deal with the fact that BigQuery does not have a concept of a schema
// TODO Deal with the fact that BigQuery does not have a concept of a database
// TODO Deal with datasets

type BigQuery struct {
	db        *sql.DB
	projectID string
	datasetID string
}

func (b BigQuery) ReturnFullPathRequirement(tableName string) string {
	return fmt.Sprintf("%s.%s.%s", b.projectID, b.datasetID, tableName)
}

func (b BigQuery) SeedsDropTableQuery(tableName string) string {
	return databaseImplementationBase.BaseForSeedsDeleteTable(tableName)
}

func (b BigQuery) SeedsCreateTableQuery(tableName string, columns []string) (string, error) {
	return databaseImplementationBase.BaseForSeedsCreateTableSpecifyingTextType("STRING", tableName, columns)
}

func (b BigQuery) SeedsInsertIntoTableQuery(tableName string, columns []string, values [][]string) (string, error) {
	return databaseImplementationBase.BaseForSeedsInsertTable(tableName, columns, values)
}

func NewBigQuery(ctx context.Context, projectID string, datasetID string) (BigQuery, error) {
	db, err := sql.Open("bigquery",
		fmt.Sprintf("bigquery://%s/%s", projectID, datasetID))
	if err != nil {
		log.Fatal(err)
	}
	return BigQuery{
		db:        db,
		projectID: projectID,
		datasetID: datasetID,
	}, nil
}

type Table struct {
	Name string `bigquery:"table_name"`
}

func (b BigQuery) ListTables(ctx context.Context) ([]string, error) {
	rows, err := b.db.QueryContext(ctx, "SELECT table_name FROM INFORMATION_SCHEMA.TABLES")
	if err != nil {
		return nil, err
	}

	var tables []string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}

	return tables, nil
}

func (b BigQuery) ListViews(ctx context.Context) ([]string, error) {
	rows, err := b.db.QueryContext(ctx, "SELECT table_name FROM INFORMATION_SCHEMA.VIEWS")
	if err != nil {
		return nil, err
	}

	var views []string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		views = append(views, table)
	}

	return views, nil
}

func (b BigQuery) ListColumns(ctx context.Context, table string) ([]string, error) {
	rows, err := b.db.QueryContext(ctx, `SELECT column_name FROM INFORMATION_SCHEMA.COLUMNS WHERE table_name = $1`, table)
	if err != nil {
		return nil, err
	}

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

func (b BigQuery) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return b.db.ExecContext(ctx, query, args...)
}

func (b BigQuery) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return b.db.QueryContext(ctx, query, args...)
}

func (b BigQuery) Close(ctx context.Context) error {
	return b.db.Close()
}
