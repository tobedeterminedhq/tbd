// Package databaseImplementationBase provides a base class for the database implementations.
package databaseImplementationBase

import (
	"bytes"
	"fmt"
	"text/template"
)

// TODO Over time can swap returning strings to returning buffers.
// TODO Over time we will want to separate this out for each database type from the calls to them. So that the sql can be rendered wherever without having a need for a database connection.

func BaseForSeedsDeleteTable(tableName string) string {
	return "DROP TABLE IF EXISTS " + tableName
}

func BaseForSeedsCreateTable(tableName string, columns []string) (string, error) {
	return BaseForSeedsCreateTableSpecifyingTextType("TEXT", tableName, columns)
}

func BaseForSeedsCreateTableSpecifyingTextType(textType, tableName string, columns []string) (string, error) {
	buf := new(bytes.Buffer)
	if err := seedsCreateStatement.Execute(buf, sqlTemplateType{
		TableName: tableName,
		Columns:   columns,
		TextType:  textType,
	}); err != nil {
		return "", fmt.Errorf("executing template: %w", err)
	}
	return buf.String(), nil
}

func BaseForSeedsInsertTable(tableName string, columns []string, values [][]string) (string, error) {
	buf := new(bytes.Buffer)
	if err := seedsTemplate.Execute(buf, sqlTemplateType{
		TableName: tableName,
		Columns:   columns,
		Values:    values,
	}); err != nil {
		return "", fmt.Errorf("executing template: %w", err)
	}
	return buf.String(), nil
}

type sqlTemplateType struct {
	TableName string
	TextType  string
	Columns   []string
	Values    [][]string
}

var seedsTemplate = template.Must(template.New("sql").Parse(`INSERT INTO {{.TableName}} ({{range $i, $col := .Columns}}{{if $i}},{{end}}{{$col}}{{end}}) VALUES {{range $i, $row := .Values}}{{if $i}},{{end}}({{range $j, $col := $row}}{{if $j}},{{end}}'{{$col}}'{{end}}){{end}}`))

var seedsCreateStatement = template.Must(template.New("sql").Parse(`CREATE TABLE {{.TableName}} ({{range $i, $col := .Columns}}{{if $i}},{{end}}{{$col}} {{$.TextType}}{{end}})`))
