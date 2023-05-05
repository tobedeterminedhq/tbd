package lib

import (
	"bytes"
	"io"
	"strings"
	"text/template"

	servicev1 "github.com/tobedeterminedhq/tbd/proto_gen/go/tbd/service/v1"
)

func renderColumnTestFromString(templateString string, path string, columnName string, customInfo map[string]string) (string, error) {
	const name = "custom_test"
	t := template.New(name)
	t, err := t.Parse(templateString)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	err = renderColumnTestWithCustomInfo(t, buf, path, columnName, customInfo)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func renderColumnTestWithCustomInfo(t *template.Template, w io.Writer, path string, columnName string, customInfo map[string]string) error {
	m := make(map[string]string, 2+len(customInfo))
	m["Model"] = path
	m["Column"] = columnName
	for k, v := range customInfo {
		m[strings.ToTitle(k[0:1])+k[1:]] = v
	}
	return t.Execute(w, m)
}

func GenerateTestNameCustomColumn(test *servicev1.TestCustomColumn) string {
	return "test_" + test.GetModel() + "_" + test.GetColumn() + "_" + test.GetTestName()
}
