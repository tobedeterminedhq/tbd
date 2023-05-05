package lib

import (
	"bytes"
	"io"
)

// ParseModelSchemasToViews takes in a reader and reads it to a View file
// nameReplacingStrategy takes in the reference name and replaces it with whatever strategy is necessary.
func ParseModelSchemasToViews(
	fileReader io.Reader,
	viewName string,
	configSchemaName string,
	nameReplacingStrategy func(name string) string,
) (string, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(fileReader)
	if err != nil {
		return "", err
	}
	originalSelectStatement := buf.String()
	referenceSearch, err := returnReferenceSearch(configSchemaName)
	if err != nil {
		return "", err
	}
	outSelect := referenceSearch.ReplaceAllStringFunc(originalSelectStatement, nameReplacingStrategy)
	return returnSQLModelTemplate(viewName, outSelect), nil
}

func returnSQLModelTemplate(name string, selectStatement string) string {
	return `DROP VIEW IF EXISTS ` + name + `; CREATE VIEW ` + name + ` AS ` + selectStatement + `;`
}
