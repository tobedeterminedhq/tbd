package lib

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/fs"
	"sort"
	"strings"

	"github.com/tobedeterminedhq/tbd/lib/databases"

	"github.com/samber/lo"
	servicev1 "github.com/tobedeterminedhq/tbd/proto_gen/go/tbd/service/v1"
)

// ProjectAndFsToSqlForViews returns both the sql for the seeds and the sql for the models.
// The returned tuple is a slice of [2]string where the first element is the name of the seed or model and the second
// element is the sql. The seeds are returned first.
//
// The returned Sql is in the shape of a view for models and tables for seeds.
//
// The onlyModels flag is used to only return the sql for the models.
func ProjectAndFsToSqlForViews(p *servicev1.Project, fsfs fs.FS, database databases.Database, onlyModels bool, doNotIncludeSeedsData bool) ([][2]string, error) {
	seeds := lo.Values(p.Seeds)
	sort.Slice(seeds, func(i, j int) bool {
		return seeds[i].Name < seeds[j].Name
	})
	outSeeds := make([][2]string, len(seeds))
	for i, s := range seeds {
		bs, err := fs.ReadFile(fsfs, s.GetFilePath())
		if err != nil {
			return nil, err
		}
		reader := strings.NewReader(string(bs))
		out, err := ParseTableSchemaSeeds(database, s.GetName(), reader, doNotIncludeSeedsData)
		if err != nil {
			return nil, err
		}
		outSeeds[i] = [2]string{s.GetName(), strings.Join(out, ";")}
	}

	graph, err := ProjectToGraph(p)
	if err != nil {
		return nil, err
	}

	sorted, err := graph.graph.GetNodeSorted()
	if err != nil {
		return nil, err
	}
	outModels := make([][2]string, len(p.Models))
	i := 0
	for _, node := range sorted {
		if _, ok := p.Seeds[node]; ok {
			continue
		}
		if _, okTest := p.Tests[node]; okTest {
			continue
		}

		if _, okSource := p.Sources[node]; okSource {
			continue
		}
		m, okModel := p.Models[node]
		if !okModel {
			models := lo.Keys(p.Models)
			sort.Strings(models)
			return nil, fmt.Errorf("model %s not found in models %v", node, models)
		}

		// TODO Make even lazier without reading early
		bs, err := fs.ReadFile(fsfs, m.GetFilePath())
		if err != nil {
			return nil, err
		}
		out, err := ParseModelSchemasToViews(
			bytes.NewReader(bs),
			node,
			p.GetConfig().GetSchemaName(),
			func(s string) string {
				replaced := ReplaceReferenceStringFound(p.GetConfig().GetSchemaName(), p.GetSources())(s)
				replaced = strings.Trim(replaced, " ")
				return " " + database.ReturnFullPathRequirement(replaced)
			},
		)
		if err != nil {
			return nil, err
		}

		outModels[i] = [2]string{node, out}
		i++
	}

	if onlyModels {
		return outModels, nil
	}
	return append(outSeeds, outModels...), nil
}

// ReturnSQLForModel returns the sql for a model but does so in the shape to create a view
func ReturnSQLForModel(p *servicev1.Project, database databases.Database, fsfs fs.FS, modelName string) (string, error) {
	output, err := ProjectAndFsToSqlForViews(p, fsfs, database, true, false)
	if err != nil {
		return "", err
	}
	for _, o := range output {
		if o[0] == modelName {
			return o[1], nil
		}
	}
	return "", fmt.Errorf("model %s not found in project", modelName)
}

// ProjectAndFsToQuerySql returns a SELECT statement for the model specified such that it can be used.
func ProjectAndFsToQuerySql(p *servicev1.Project, fsfs fs.FS, model string) (string, error) {
	_, okSeed := p.Seeds[model]
	_, okModel := p.Models[model]
	if !okSeed && !okModel {
		return "", fmt.Errorf("model %s not found in models %v nor seeds %v", model, lo.Keys(p.Models), lo.Keys(p.Seeds))
	}
	graph, err := ProjectToGraph(p)
	if err != nil {
		return "", err
	}

	models, err := graph.graph.ReturnSubGraphNodes(model)
	if err != nil {
		return "", err
	}

	nodeNames := make([]string, len(models))
	for i, m := range models {
		nodeNames[i], err = graph.graph.GetNodeName(m)
		if err != nil {
			return "", err
		}
	}
	toProcess := make([]nodeWithName, len(nodeNames))
	for i, m := range nodeNames {
		seed, isSeed := p.GetSeeds()[m]
		model, isModel := p.GetModels()[m]
		switch {
		case isSeed && isModel:
			return "", fmt.Errorf("node %s is both a seed and a model", m)
		case !isSeed && !isModel:
			return "", fmt.Errorf("node %s is neither a seed or a model", m)
		case isSeed:
			toProcess[i] = nodeWithName{node: seed, name: m}
		case isModel:
			toProcess[i] = nodeWithName{node: model, name: m}
		default:
			return "", fmt.Errorf("node %s is neither a seed or a model", m)
		}
	}
	return convertToSelectStatements(p.GetConfig(), fsfs, p.GetSources(), toProcess)
}

type nodeWithName struct {
	name string
	node interface{}
}

// convertToSelectStatements takes in an array of models and returns a string that can be used in a select statement.
// It also replaces any tbd.references with the actual name that is in the select. It uses no views.
// TODO Need to figure out in the future how to handle the case of seeds.
//
// array of models is in the shape of [][2]string where the first element is the name of the model and the second element is the sql
func convertToSelectStatements(config *servicev1.Configuration, fsfs fs.FS, sources map[string]*servicev1.Source, values []nodeWithName) (string, error) {
	selects := make([][2]string, len(values))
	for i, v := range values {
		switch v := v.node.(type) {
		case *servicev1.Source:
			out := renderSourceSelectStatement(v)
			selects[i] = [2]string{v.GetName(), out}
		case *servicev1.Seed:
			out, err := renderSeedSelectStatement(fsfs, v)
			if err != nil {
				return "", err
			}
			selects[i] = [2]string{v.GetName(), out}
		case *servicev1.Model:
			out, err := renderModelSelect(fsfs, config, v, sources)
			if err != nil {
				return "", err
			}
			selects[i] = [2]string{v.GetName(), out}
		default:
			return "", fmt.Errorf("unknown type %T", v)
		}
	}
	switch len(selects) {
	case 0:
		return "", fmt.Errorf("array of models to select is empty")
	case 1:
		return selects[0][1], nil
	case 2:
		out := fmt.Sprintf("WITH %s AS (%s) %s", selects[0][0], selects[0][1], selects[1][1])
		return out, nil
	default:
		out := "WITH\n"
		for _, a := range selects[0 : len(selects)-1] {
			out += fmt.Sprintf("%s AS (%s),\n ", a[0], a[1])
		}
		out = out[:len(out)-3] + " " + "SELECT * FROM (" + selects[len(selects)-1][1] + ")"
		return out, nil
	}
}

func renderSeedSelectStatement(fsfs fs.FS, seed *servicev1.Seed) (string, error) {
	bs, err := fs.ReadFile(fsfs, seed.GetFilePath())
	if err != nil {
		return "", err
	}
	r := csv.NewReader(bytes.NewReader(bs))
	records, err := r.ReadAll()
	if err != nil {
		return "", err
	}
	return renderSeedSelectStatementString(records[0], records[1:]), nil
}

func renderSeedSelectStatementString(headers []string, values [][]string) string {
	var headersPart string
	for i, h := range headers {
		if i == len(headers)-1 {
			headersPart += fmt.Sprintf("column%d AS %s", i+1, h)
			break
		}
		headersPart += fmt.Sprintf("column%d AS %s, ", i+1, h)
	}
	var valuesPart string
	for i, v := range values {
		var valuePart string
		for j, vv := range v {
			if j == len(v)-1 {
				valuePart += fmt.Sprintf("'%s'", vv)
				break
			}
			valuePart += fmt.Sprintf("'%s', ", vv)
		}
		if i == len(values)-1 {
			valuesPart += fmt.Sprintf("(%s)", valuePart)
			break
		}
		valuesPart += fmt.Sprintf("(%s), ", valuePart)
	}
	return "SELECT " + headersPart + " FROM (VALUES " + valuesPart + ")"
}

func renderSourceSelectStatement(source *servicev1.Source) string {
	return fmt.Sprintf("SELECT * FROM %s", source.GetPath())
}

func renderModelSelect(fsfs fs.FS, config *servicev1.Configuration, model *servicev1.Model, sources map[string]*servicev1.Source) (string, error) {
	bs, err := fs.ReadFile(fsfs, model.GetFilePath())
	if err != nil {
		return "", err
	}
	out := string(bs)
	referenceSearch, err := returnReferenceSearch(config.GetSchemaName())
	if err != nil {
		return "", err
	}
	replaced := referenceSearch.ReplaceAllStringFunc(out, ReplaceReferenceStringFound(config.GetSchemaName(), sources))
	return replaced, nil
}

// ReplaceReferenceStringFound takes in a reference such as `tbd.raw_orders` where the schemaName is specified as `tbd`
// and replaces it with the value of the reference raw_orders. This is used to replace references in the sql files with
// the actual sql
//
// For a reference to be replaced it must be in the form of `schemaName.referenceName`
//
// For references in general, the schemaName is just removed. In the case that the reference is a source. The string is
// replaced by the path property of the source.
func ReplaceReferenceStringFound(schemaName string, sources map[string]*servicev1.Source) func(string) string {
	return func(s string) string {
		model := s[len(schemaName)+2:]
		if source, ok := sources[model]; ok {
			return " " + source.GetPath()
		}
		return " " + s[len(schemaName)+2:]
	}
}

// ReturnTestsSQL returns sql tests to run in no order but with the name pointing to the test
// TODO Need to write a test for this
// TODO Need to add safe adders for the tests map
func ReturnTestsSQL(p *servicev1.Project, fsfs fs.FS) (map[string]string, error) {
	tests := make(map[string]string, len(p.Tests))
	referenceSearch, err := returnReferenceSearch(p.GetConfig().GetSchemaName())
	if err != nil {
		return nil, err
	}
	for _, t := range p.Tests {
		switch {
		case t.GetSql() != nil:
			test := t.GetSql()
			bs, err := fs.ReadFile(fsfs, test.GetFilePath())
			if err != nil {
				return nil, err
			}
			out := referenceSearch.ReplaceAllStringFunc(string(bs), ReplaceReferenceStringFound(p.GetConfig().GetSchemaName(), p.GetSources()))
			tests[test.GetName()] = out
		case t.GetUnique() != nil:
			test := t.GetUnique()
			out := GenerateTestSqlUnique(test)
			name := GenerateTestNameUnique(test)
			tests[name] = out
		case t.GetNotNull() != nil:
			test := t.GetNotNull()
			out := GenerateTestSqlNotNull(test)
			name := GenerateTestNameNotNull(test)
			tests[name] = out
		case t.GetRelationship() != nil:
			test := t.GetRelationship()
			out := GenerateTestSqlRelationship(test)
			name := GenerateTestNameRelationship(test)
			tests[name] = out
		case t.GetAcceptedValues() != nil:
			test := t.GetAcceptedValues()
			out := GenerateTestSqlAcceptedValues(test)
			name := GenerateTestNameAcceptedValues(test)
			tests[name] = out
		case t.GetCustomColumn() != nil:
			test := t.GetCustomColumn()
			name := GenerateTestNameCustomColumn(test)
			tests[name] = test.GetRenderedSql()
		default:
			return nil, fmt.Errorf("test '%v' is not a sql/not_null/unique/relationship/custom test", t)
		}
	}
	return tests, nil
}
