package lib

import (
	"fmt"

	servicev1 "github.com/benfdking/tbd/proto/gen/go/tbd/service/v1"
)

// GenerateTestSqlNotNull generates a SQL test that checks that the given column does not contain any nulls.
//
// The test was generated with the following template but has been optimized
// template.Must(template.New("sqlNotNull").Parse(`SELECT * FROM {{.Table}} WHERE {{.Column}} IS NULL`))
func GenerateTestSqlNotNull(test *servicev1.TestNotNull) string {
	return "SELECT * FROM " + test.GetPath() + " WHERE " + test.GetColumn() + " IS NULL"
}

func GenerateTestNameNotNull(test *servicev1.TestNotNull) string {
	return "test_" + test.GetModel() + "_" + test.GetColumn() + "_not_null"
}

func GenerateTestSqlUnique(test *servicev1.TestUnique) string {
	return fmt.Sprintf(`SELECT * FROM (
    SELECT %s
    FROM %s
    WHERE %s IS NOT NULL
    GROUP BY %s
    HAVING count(*) > 1
)`, test.GetColumn(), test.GetPath(), test.GetColumn(), test.GetColumn())
}

func GenerateTestNameUnique(test *servicev1.TestUnique) string {
	return "test_" + test.GetModel() + "_" + test.GetColumn() + "_unique"
}

func GenerateTestSqlRelationship(test *servicev1.TestRelationship) string {
	return fmt.Sprintf(
		`SELECT * FROM %s WHERE %s IS NOT NULL AND %s NOT IN (SELECT %s FROM %s)`,
		test.GetSourcePath(),
		test.GetSourceColumn(),
		test.GetSourceColumn(),
		test.GetTargetColumn(),
		test.GetTargetPath(),
	)
}

func GenerateTestNameRelationship(test *servicev1.TestRelationship) string {
	return fmt.Sprintf(
		"test_%s_%s_relationship_%s_%s",
		test.GetSourceModel(),
		test.GetSourceColumn(),
		test.GetTargetModel(),
		test.GetTargetColumn(),
	)
}

func GenerateTestNameAcceptedValues(test *servicev1.TestAcceptedValues) string {
	return fmt.Sprintf(
		"test_%s_%s_accepted_values",
		test.GetModel(),
		test.GetColumn(),
	)
}

func GenerateTestSqlAcceptedValues(test *servicev1.TestAcceptedValues) string {
	return fmt.Sprintf(
		`SELECT * FROM %s WHERE %s IS NOT NULL AND %s NOT IN (%s)`,
		test.GetPath(),
		test.GetColumn(),
		test.GetColumn(),
		generateSqlInList(test.GetAcceptedValues()),
	)
}

func generateSqlInList(values []string) string {
	var inList string
	for _, value := range values {
		inList += "'" + value + "',"
	}
	return inList[:len(inList)-1]
}

type StandardTestTypes string

const (
	StandardTestTypeSqlNotNull     StandardTestTypes = "not_null"
	StandardTestTypeSqlUnique      StandardTestTypes = "unique"
	StandardTestTypeRelationship   StandardTestTypes = "relationship"
	StandardTestTypeAcceptedValues StandardTestTypes = "accepted_values"
)
