package lib

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	servicev1 "github.com/benfdking/tbd/proto/gen/go/tbd/service/v1"
	"github.com/samber/lo"
	"sigs.k8s.io/yaml"
)

// ParseConfigFromPath runs ParseConfig given path.
func ParseConfigFromPath(path string) (*servicev1.Configuration, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}
	defer f.Close()
	return ParseConfig(f)
}

// ParseConfig parses bytes to configuration.  It also applies any defaults to the configuration, as well as validates the
// configuration.
func ParseConfig(r io.Reader) (*servicev1.Configuration, error) {
	bs, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("reading config: %w", err)
	}
	var c servicev1.Configuration
	if err := yaml.Unmarshal(bs, &c); err != nil {
		return nil, fmt.Errorf("umarshaling yaml: %w", err)
	}
	// apply defaults
	withDefs := applyDefaultsToConfig(&c)
	// validate
	if err := validateConfig(withDefs); err != nil {
		return nil, fmt.Errorf("validating config: %w", err)
	}
	return withDefs, nil
}

// ValidateModelName validates a model name.
func ValidateModelName(name string) error {
	if !validateConfigSchemaName.MatchString(name) {
		return fmt.Errorf("model name must match %s", validateConfigSchemaName.String())
	}
	return nil
}

var validateConfigSchemaName = regexp.MustCompile(`^[a-z][a-z0-9_]*$`)

func validateConfig(c *servicev1.Configuration) error {
	if !validateConfigSchemaName.MatchString(c.GetSchemaName()) {
		return fmt.Errorf("schema name must match %s", validateConfigSchemaName.String())
	}
	return nil
}

func applyDefaultsToConfig(c *servicev1.Configuration) *servicev1.Configuration {
	name := "tbd"
	if c.SchemaName == nil {
		c.SchemaName = &name
	}
	if *c.SchemaName == "" {
		c.SchemaName = &name
	}
	return c
}

// ParseProject parses a whole project into a project object
func ParseProject(c *servicev1.Configuration, fs fs.FS, projectRoot string) (*servicev1.Project, error) {
	seeds, err := parseSeeds(fs, projectRoot, c)
	if err != nil {
		return nil, fmt.Errorf("parsing seeds: %w", err)
	}

	// parsing project files
	projectFiles, err := parseProjectFiles(fs, projectRoot, c)
	if err != nil {
		return nil, fmt.Errorf("parsing project files: %w", err)
	}
	sources := parseSources(projectFiles)

	customTests, err := parseCustomTests(fs, projectRoot, c)
	if err != nil {
		return nil, fmt.Errorf("parsing custom tests: %w", err)
	}

	allModelDefinitions := lo.Values(projectFiles)
	fileModels := lo.FlatMap(allModelDefinitions, func(v *servicev1.ProjectFile, i int) []*servicev1.ProjectFile_Model {
		return v.GetModels()
	})
	m := ModelDefinitions(lo.KeyBy(fileModels, func(v *servicev1.ProjectFile_Model) string {
		return v.GetName()
	}))

	models, err := parseModels(fs, m, projectRoot, c)
	if err != nil {
		return nil, fmt.Errorf("parsing models: %w", err)
	}

	pathMap := createPathMap(models, sources)
	tests, err := parseTests(c, fs, projectRoot, pathMap, customTests, projectFiles)
	if err != nil {
		return nil, err
	}

	return &servicev1.Project{
		Config:      c,
		CustomTests: customTests,
		Seeds:       seeds,
		// TODO: Add sources
		Models:       models,
		ProjectFiles: projectFiles,
		Tests:        tests,
		Sources:      sources,
	}, nil
}

func parseSources(projectFiles map[string]*servicev1.ProjectFile) map[string]*servicev1.Source {
	sourceDefinitions := lo.FlatMap(lo.Entries(projectFiles), func(v lo.Entry[string, *servicev1.ProjectFile], i int) []*servicev1.Source {
		return lo.Map(v.Value.GetSources(), func(s *servicev1.ProjectFile_Source, i int) *servicev1.Source {
			return &servicev1.Source{
				Name:        s.GetName(),
				Description: s.GetDescription(),
				Path:        s.GetPath(),
				Columns:     nil,
				FilePath:    v.Key,
			}
		})
	})
	return lo.KeyBy(sourceDefinitions, func(v *servicev1.Source) string {
		return v.GetName()
	})
}

func parseTests(
	c *servicev1.Configuration,
	fs fs.FS,
	projectRoot string,
	pathMap PathMap,
	customTests map[string]*servicev1.CustomTest,
	projectFiles map[string]*servicev1.ProjectFile,
) (map[string]*servicev1.Test, error) {
	tests, err := parseSQLTests(fs, projectRoot, c)
	if err != nil {
		return nil, fmt.Errorf("parsing sql tests: %w", err)
	}

	columnTests, err := parseColumnTests(customTests, projectFiles, pathMap)
	if err != nil {
		return nil, err
	}

	for k, v := range columnTests {
		err := errorMapAdder(tests, k, v)
		if err != nil {
			return nil, err
		}
	}
	return tests, nil
}

// TODO Add regex testing as well
// TODO Need to be able to add information about columns to seeds

func parseSeeds(filesystem fs.FS, projectRoot string, c *servicev1.Configuration) (map[string]*servicev1.Seed, error) {
	allPaths := make(map[string]struct{})
	for _, path := range c.GetSeedPaths() {

		files, err := listAllFilesInDirectoryRecursively(filesystem, filepath.Join(projectRoot, path))
		if err != nil {
			return nil, fmt.Errorf("listing all files for path %s: %w", path, err)
		}
		filtered := lo.Filter[string](files, byType(csvFileType))
		for _, v := range filtered {
			allPaths[v] = struct{}{}
		}
	}
	seeds := make(map[string]*servicev1.Seed, len(allPaths))
	for path := range allPaths {
		_, file := filepath.Split(path)
		name := strings.TrimSuffix(file, filepath.Ext(path))
		seeds[name] = &servicev1.Seed{
			Name:     name,
			FilePath: path,
		}
	}
	return seeds, nil
}

// TODO Need to parse tests
//	Parse Yml tests and sql tests fully

// parseModels parses models from the .sql file without adding information from the yml file
func parseModels(filesystem fs.FS, modelDefinition ModelDefinitions, projectRoot string, c *servicev1.Configuration) (map[string]*servicev1.Model, error) {
	allPaths := make(map[string]struct{})
	for _, path := range c.GetModelPaths() {
		files, err := listAllFilesInDirectoryRecursively(filesystem, filepath.Join(projectRoot, path))
		if err != nil {
			return nil, err
		}
		filteredFiles := lo.Filter(files, byType(sqlFileType))
		for _, v := range filteredFiles {
			allPaths[v] = struct{}{}
		}
	}

	models := make([]*servicev1.Model, len(allPaths))
	i := 0
	for path := range allPaths {
		var err error
		models[i], err = parseModel(filesystem, modelDefinition, c, path)
		if err != nil {
			return nil, err
		}
		i++
	}

	// convert to map
	outs := make(map[string]*servicev1.Model, len(models))
	for _, model := range models {
		err := errorMapAdder(outs, model.GetName(), model)
		if err != nil {
			return nil, err
		}
	}
	return outs, nil
}

type ModelDefinitions map[string]*servicev1.ProjectFile_Model

// parseModel parses a model from the .sql file and does not add the information from the yml file.
func parseModel(fsfs fs.FS, modelDefinitions ModelDefinitions, c *servicev1.Configuration, path string) (*servicev1.Model, error) {
	_, file := filepath.Split(path)
	name := strings.TrimSuffix(file, filepath.Ext(path))
	var description string
	if v, ok := modelDefinitions[name]; ok {
		description = v.GetDescription()
	}
	// TODO Need to pass in description and columns
	model := &servicev1.Model{
		Name:        name,
		Description: description,
		FilePath:    path,
	}
	referenceSearch, err := returnReferenceSearch(c.GetSchemaName())
	if err != nil {
		return nil, err
	}
	bs, err := fs.ReadFile(fsfs, model.GetFilePath())
	if err != nil {
		return nil, err
	}
	all := referenceSearch.FindAllStringSubmatch(string(bs), -1)
	model.References = parseReferences(all)
	return model, nil
}

// parseReferences parses references
func parseReferences(refs [][]string) []string {
	out := make([]string, len(refs))
	for i := range refs {
		out[i] = refs[i][1]
	}
	return out
}

// parseProjectFiles will parse the project files that are in the models paths that will return a map of the project files
// location relative to fs.FS.
func parseProjectFiles(fsfs fs.FS, projectRoot string, c *servicev1.Configuration) (map[string]*servicev1.ProjectFile, error) {
	allPaths := make(map[string]struct{})
	for _, path := range c.GetModelPaths() {
		files, err := listAllFilesInDirectoryRecursively(fsfs, filepath.Join(projectRoot, path))
		if err != nil {
			return nil, err
		}
		filteredFiles := lo.Filter(files, byType(yml))
		for _, v := range filteredFiles {
			allPaths[v] = struct{}{}
		}
	}

	outs := make(map[string]*servicev1.ProjectFile, len(allPaths))
	for path := range allPaths {
		file, err := fsfs.Open(path)
		if err != nil {
			return nil, err
		}
		projectFile, err := ParseProjectFile(file)
		if err != nil {
			return nil, err
		}
		err = file.Close()
		if err != nil {
			return nil, err
		}
		outs[path] = projectFile
	}
	return outs, nil
}

func returnReferenceSearch(schemaName string) (*regexp.Regexp, error) {
	return regexp.Compile(`\s` + schemaName + "." + `([a-zA-Z][a-z_A-Z0-9]*)`)
}

// PathMap is a map of names to paths that includes models and sources to be used to convert model names to paths in
// test creation
type PathMap map[string]string

// TODO in the future probably want to be able to set a custom path for a model
// TODO in the future this can probably be cleaned up by using this map more widely
func createPathMap(models map[string]*servicev1.Model, sources map[string]*servicev1.Source) PathMap {
	m := make(PathMap, len(models)+len(sources))
	for k := range models {
		m[k] = k
	}
	for k, v := range sources {
		m[k] = v.GetPath()
	}
	return m
}

// parseSQLTests parses sql tests
func parseSQLTests(filesystem fs.FS, projectRoot string, c *servicev1.Configuration) (map[string]*servicev1.Test, error) {
	var allPaths []string
	for _, path := range c.GetTestPaths() {
		files, err := listAllFilesInDirectoryRecursively(filesystem, filepath.Join(projectRoot, path))
		if err != nil {
			return nil, err
		}
		allPaths = append(allPaths, lo.Filter[string](files, byType(sqlFileType))...)
	}
	referenceSearch, err := returnReferenceSearch(c.GetSchemaName())
	if err != nil {
		return nil, err
	}

	outs := make(map[string]*servicev1.Test)
	for _, path := range allPaths {
		bs, err := fs.ReadFile(filesystem, path)
		if err != nil {
			return nil, err
		}

		// TODO parsing out references

		_, file := filepath.Split(path)
		testName := "test_sql__" + strings.TrimSuffix(file, filepath.Ext(path))
		all := referenceSearch.FindAllStringSubmatch(string(bs), -1)
		sqlTest := &servicev1.Test_Sql{Sql: &servicev1.TestSQLFile{
			FilePath:   path,
			Name:       testName,
			References: parseReferences(all),
		}}

		outs[testName] = &servicev1.Test{TestType: sqlTest}
	}

	return outs, nil
}

func parseColumnTests(
	customTests map[string]*servicev1.CustomTest,
	projectFiles map[string]*servicev1.ProjectFile,
	pathMap PathMap,
) (map[string]*servicev1.Test, error) {
	outs := make(map[string]*servicev1.Test)

	for filePath, projectFile := range projectFiles {
		for _, source := range projectFile.Sources {
			for _, column := range source.Columns {
				m, err := parseColumTestsForModelOrSource(customTests, column, pathMap, filePath, source.GetName(), source.GetPath())
				if err != nil {
					return m, err
				}
				for k, v := range m {
					err := errorMapAdder(outs, k, v)
					if err != nil {
						return nil, err
					}
				}
			}
		}
		for _, model := range projectFile.Models {
			for _, column := range model.Columns {
				m, err := parseColumTestsForModelOrSource(customTests, column, pathMap, filePath, model.GetName(), model.GetName())
				if err != nil {
					return m, err
				}
				for k, v := range m {
					err := errorMapAdder(outs, k, v)
					if err != nil {
						return nil, err
					}
				}
			}
		}
	}
	return outs, nil
}

func parseColumTestsForModelOrSource(
	customTests map[string]*servicev1.CustomTest,
	column *servicev1.ProjectFile_Column,
	pathMap PathMap,
	filePath string,
	modelName string,
	modelPath string,
) (map[string]*servicev1.Test, error) {
	outs := make(map[string]*servicev1.Test)
	{
		for _, test := range column.Tests {
			switch StandardTestTypes(test.GetType()) {
			case StandardTestTypeSqlNotNull:
				out := &servicev1.TestNotNull{
					FilePath: filePath,
					Model:    modelName,
					Path:     modelPath,
					Column:   column.GetName(),
				}
				outs[GenerateTestNameNotNull(out)] = &servicev1.Test{
					TestType: &servicev1.Test_NotNull{NotNull: out},
				}
			case StandardTestTypeSqlUnique:
				out := &servicev1.TestUnique{
					FilePath: filePath,
					Model:    modelName,
					Path:     modelPath,
					Column:   column.GetName(),
				}
				outs[GenerateTestNameUnique(out)] = &servicev1.Test{
					TestType: &servicev1.Test_Unique{Unique: out},
				}
			case StandardTestTypeRelationship:
				// TODO Refactor this to a map of functions
				info := test.GetInfo()
				target, ok := info["model"]
				if !ok {
					return nil, fmt.Errorf("'model' must be specified in column test information")
				}
				targetColumn, ok := info["column"]
				if !ok {
					return nil, fmt.Errorf("'column' must be specified in column test information")
				}
				targetPath, ok := pathMap[target]
				if !ok {
					return nil, fmt.Errorf("could not find %s in path map %v", target, lo.Keys(pathMap))
				}
				out := &servicev1.TestRelationship{
					FilePath:     filePath,
					SourceModel:  modelName,
					SourcePath:   modelPath,
					SourceColumn: column.GetName(),
					// TODO This will need to be changed so that a target model can be a source model and the
					//   target column path can be right
					TargetModel:  target,
					TargetPath:   targetPath,
					TargetColumn: targetColumn,
				}
				outs[GenerateTestNameRelationship(out)] = &servicev1.Test{
					TestType: &servicev1.Test_Relationship{Relationship: out},
				}
			case StandardTestTypeAcceptedValues:
				info := test.GetInfo()
				acceptableValues, ok := info["values"]
				if !ok {
					return nil, fmt.Errorf("'values' must be specified in column test information")
				}
				out := &servicev1.TestAcceptedValues{
					FilePath:       filePath,
					Model:          modelName,
					Path:           modelPath,
					Column:         column.GetName(),
					AcceptedValues: strings.Split(acceptableValues, ","),
				}
				outs[GenerateTestNameAcceptedValues(out)] = &servicev1.Test{
					TestType: &servicev1.Test_AcceptedValues{AcceptedValues: out},
				}
			default:
				name, out, err := parseCustomColumnTest(customTests, test.GetType(), modelName, modelPath, column.GetName(), test.GetInfo())
				if err != nil {
					return nil, err
				}
				outs[name] = out
			}
		}
	}
	return outs, nil
}

func parseCustomTests(filesystem fs.FS, projectRoot string, c *servicev1.Configuration) (map[string]*servicev1.CustomTest, error) {
	var allPaths []string
	for _, path := range c.GetCustomTestPaths() {
		files, err := listAllFilesInDirectoryRecursively(filesystem, filepath.Join(projectRoot, path))
		if err != nil {
			return nil, err
		}
		allPaths = append(allPaths, lo.Filter[string](files, byType(sqlFileType))...)
	}

	outs := make(map[string]*servicev1.CustomTest)
	for _, path := range allPaths {
		name := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
		file, err := filesystem.Open(path)
		if err != nil {
			return nil, fmt.Errorf("failed to open file %s: %w", path, err)
		}
		defer file.Close()
		contents, err := io.ReadAll(file)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %w", path, err)
		}
		out := &servicev1.CustomTest{
			FilePath: path,
			Name:     name,
			Sql:      string(contents),
		}
		err = errorMapAdder(outs, name, out)
		if err != nil {
			return nil, err
		}
	}
	return outs, nil
}

func parseCustomColumnTest(
	customTests map[string]*servicev1.CustomTest,
	testName string,
	model string,
	path string,
	column string,
	info map[string]string,
) (
	name string,
	out *servicev1.Test,
	err error,
) {
	test, ok := customTests[testName]
	if !ok {
		return "", nil, fmt.Errorf("test type %s not found", testName)
	}
	rendered, err := renderColumnTestFromString(test.GetSql(), path, column, info)
	if err != nil {
		return "", nil, err
	}
	outTest := &servicev1.TestCustomColumn{
		TestFilePath: test.GetFilePath(),
		TestName:     testName,
		OriginalSql:  test.GetSql(),
		Model:        model,
		Path:         path,
		Column:       column,
		Info:         info,
		RenderedSql:  rendered,
	}
	name = GenerateTestNameCustomColumn(outTest)
	return name,
		&servicev1.Test{
			TestType: &servicev1.Test_CustomColumn{
				CustomColumn: outTest,
			},
		},
		nil
}
