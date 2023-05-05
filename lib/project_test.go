package lib

import (
	"fmt"
	"io/fs"
	"testing"

	"github.com/benfdking/tbd/go/lib/databasesImplementation"

	"github.com/stretchr/testify/require"

	servicev1 "github.com/benfdking/tbd/proto/gen/go/tbd/service/v1"
	"github.com/stretchr/testify/assert"
)

func TestReplaceReferenceStringFound(t *testing.T) {
	rxp, err := returnReferenceSearch("tbd")
	require.NoError(t, err, "returnReferenceSearch()")

	tests := []struct {
		name    string
		schema  string
		sources map[string]*servicev1.Source
		s       string
		want    string
	}{
		{
			name:    "simple example",
			schema:  "tbd",
			sources: map[string]*servicev1.Source{},
			s:       "FROM tbd.raw_orders",
			want:    "FROM raw_orders",
		},
		{
			name:   "simple example with reference to source",
			schema: "tbd",
			sources: map[string]*servicev1.Source{
				"raw_orders_reference": {
					Name: "raw_orders_reference",
					Path: "schema.raw_orders_that_was_translated",
				},
			},
			s:    "FROM tbd.raw_orders_reference",
			want: "FROM schema.raw_orders_that_was_translated",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, rxp.ReplaceAllStringFunc(tt.s, ReplaceReferenceStringFound(tt.schema, tt.sources)), "ReplaceReferenceStringFound(%v)", tt.s)
		})
	}
}

func Benchmark_replaceStringFound(b *testing.B) {
	in := " tbd.raw_orders"

	for n := 0; n < b.N; n++ {
		ReplaceReferenceStringFound(in, map[string]*servicev1.Source{})
	}
}

func TestReturnTestsSQL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		customTests map[string]*servicev1.CustomTest
		tests       map[string]*servicev1.Test
		want        map[string]string
		fs          fs.FS
		wantErr     assert.ErrorAssertionFunc
	}{
		{
			name: "simple example with different types",
			customTests: map[string]*servicev1.CustomTest{
				"gte": {
					FilePath: "custom_test/gte.sql",
					Name:     "gte",
					Sql:      "SELECT *\nFROM {{ .Model }} WHERE {{ .Column }} < {{ .Value }}",
				},
			},
			fs: func() fs.FS {
				f, err := NewFileSystem(&servicev1.FileSystem{})
				if err != nil {
					panic("error creating filesystem")
				}
				return f
			}(),
			tests: map[string]*servicev1.Test{
				"not_null": {
					TestType: &servicev1.Test_NotNull{
						NotNull: &servicev1.TestNotNull{
							FilePath: "test.yaml",
							Model:    "hello",
							Path:     "hello",
							Column:   "world",
						},
					},
				},
				"unique": {
					TestType: &servicev1.Test_Unique{
						Unique: &servicev1.TestUnique{
							FilePath: "test.yaml",
							Model:    "hello",
							Path:     "hello",
							Column:   "world",
						},
					},
				},
				"relationship": {
					TestType: &servicev1.Test_Relationship{
						Relationship: &servicev1.TestRelationship{
							FilePath:     "test.yaml",
							SourceModel:  "hello",
							SourcePath:   "hello",
							SourceColumn: "world",
							TargetModel:  "hello_target",
							TargetPath:   "hello_target",
							TargetColumn: "world_target",
						},
					},
				},
				"gte": {
					TestType: &servicev1.Test_CustomColumn{
						CustomColumn: &servicev1.TestCustomColumn{
							TestFilePath: "custom_tests/gte.sql",
							TestName:     "gte",
							OriginalSql:  "SELECT * FROM {{ .Model }} WHERE {{ .Column }} < {{ .Value }}",
							Model:        "hello",
							Path:         "hello123",
							Column:       "world",
							Info: map[string]string{
								"Value": "1",
							},
							RenderedSql: "SELECT * FROM hello WHERE world < 1",
						},
					},
				},
			},
			want: map[string]string{
				"test_hello_world_not_null":                               "SELECT * FROM hello WHERE world IS NULL",
				"test_hello_world_relationship_hello_target_world_target": "SELECT * FROM hello WHERE world IS NOT NULL AND world NOT IN (SELECT world_target FROM hello_target)",
				"test_hello_world_unique":                                 "SELECT * FROM (\n    SELECT world\n    FROM hello\n    WHERE world IS NOT NULL\n    GROUP BY world\n    HAVING count(*) > 1\n)",
				"test_hello_world_gte":                                    "SELECT * FROM hello WHERE world < 1",
			},
			wantErr: assert.NoError,
		},
		{
			name:        "error test",
			customTests: map[string]*servicev1.CustomTest{},
			fs: func() fs.FS {
				f, err := NewFileSystem(&servicev1.FileSystem{})
				if err != nil {
					panic("error creating filesystem")
				}
				return f
			}(),
			tests: map[string]*servicev1.Test{
				"not_null": {},
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReturnTestsSQL(&servicev1.Project{
				Tests:       tt.tests,
				CustomTests: tt.customTests,
			},
				tt.fs,
			)
			if !tt.wantErr(t, err, fmt.Sprintf("ReturnTestsSQL(Project {Tests: %v})", tt.tests)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ReturnTestsSQL(Project {Tests: %v})", tt.tests)
		})
	}
}

func TestProjectAndFsToSqlForViews(t *testing.T) {
	t.Parallel()

	returnStandardConfig := func() *servicev1.Configuration {
		return &servicev1.Configuration{
			ModelPaths:      []string{"models"},
			SeedPaths:       []string{"seeds"},
			TestPaths:       []string{"tests"},
			CustomTestPaths: []string{"custom_tests"},
		}
	}
	emptyFileSystem := func() fs.FS {
		fileSystem, err := NewFileSystem(&servicev1.FileSystem{
			Files: map[string]*servicev1.File{},
		})
		if err != nil {
			panic(err)
		}
		return fileSystem
	}

	type args struct {
		p          *servicev1.Project
		fsfs       fs.FS
		onlyModels bool
	}
	tests := []struct {
		name    string
		args    args
		want    [][2]string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "empty project",
			args: args{
				p: &servicev1.Project{
					Config: returnStandardConfig(),
				},
				fsfs:       emptyFileSystem(),
				onlyModels: false,
			},
			want:    [][2]string{},
			wantErr: assert.NoError,
		},
		{
			name: "empty project with only models",
			args: args{
				p: &servicev1.Project{
					Config: returnStandardConfig(),
				},
				fsfs:       emptyFileSystem(),
				onlyModels: true,
			},
			want:    [][2]string{},
			wantErr: assert.NoError,
		},
		// TODO Need to add more tests
		// model and seed with output of just models and with output of models and seeds
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := databasesImplementation.NewSqlLiteInMemory()
			if err != nil {
				t.Fatal(err)
			}

			got, err := ProjectAndFsToSqlForViews(tt.args.p, tt.args.fsfs, db, tt.args.onlyModels, false)
			if !tt.wantErr(t, err, fmt.Sprintf("ProjectAndFsToSqlForViews(%v, %v, %v)", tt.args.p, tt.args.fsfs, tt.args.onlyModels)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ProjectAndFsToSqlForViews(%v, %v, %v)", tt.args.p, tt.args.fsfs, tt.args.onlyModels)
		})
	}
}

func Test_convertToSelectStatements(t *testing.T) {
	t.Parallel()

	defaultConfig := applyDefaultsToConfig(&servicev1.Configuration{})

	tests := []struct {
		name    string
		config  *servicev1.Configuration
		fsMaker func() (fs.FS, error)
		arr     []nodeWithName
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "empty array",
			config: defaultConfig,
			fsMaker: func() (fs.FS, error) {
				return NewFileSystem(&servicev1.FileSystem{
					Files: map[string]*servicev1.File{},
				})
			},
			arr:     []nodeWithName{},
			want:    "",
			wantErr: assert.Error,
		},
		{
			name: "one element which is a model",
			arr: []nodeWithName{
				{"world", &servicev1.Model{
					Name:     "world",
					FilePath: "models/world.sql",
				}},
			},
			config: defaultConfig,
			fsMaker: func() (fs.FS, error) {
				return NewFileSystem(&servicev1.FileSystem{
					Files: map[string]*servicev1.File{
						"/models/world.sql": {
							Name:     "/models/world.sql",
							Contents: []byte("SELECT tbd.hello FROM world"),
						},
					},
				})
			},
			want:    "SELECT hello FROM world",
			wantErr: assert.NoError,
		},
		{
			name: "one element which is a seed",
			arr: []nodeWithName{
				{"world", &servicev1.Seed{
					Name:     "world",
					FilePath: "seeds/world.csv",
				}},
			},
			config: defaultConfig,
			fsMaker: func() (fs.FS, error) {
				return NewFileSystem(&servicev1.FileSystem{
					Files: map[string]*servicev1.File{
						"/seeds/world.csv": {
							Name:     "/seeds/world.csv",
							Contents: []byte("header_1,header_2\nvalue_1,value_2\nvalue_3,value_4"),
						},
					},
				})
			},
			want:    "SELECT column1 AS header_1, column2 AS header_2 FROM (VALUES ('value_1', 'value_2'), ('value_3', 'value_4'))",
			wantErr: assert.NoError,
		},
		{
			name: "one element which is a source",
			arr: []nodeWithName{
				{"world", &servicev1.Source{
					Name:     "world",
					Path:     "with_schema.world_123",
					FilePath: "models/world.yaml",
				}},
			},
			config: defaultConfig,
			fsMaker: func() (fs.FS, error) {
				return NewFileSystem(&servicev1.FileSystem{
					Files: map[string]*servicev1.File{},
				})
			},
			want:    "SELECT * FROM with_schema.world_123",
			wantErr: assert.NoError,
		},
		{
			name:   "array with 4 elements, two seeds, two models",
			config: defaultConfig,
			fsMaker: func() (fs.FS, error) {
				return NewFileSystem(&servicev1.FileSystem{
					Files: map[string]*servicev1.File{
						"/seeds/world_1.csv": {
							Name:     "/seeds/world_1.csv",
							Contents: []byte("header_1,header_2\nvalue_1,value_2\nvalue_3,value_4"),
						},
						"/seeds/world_2.csv": {
							Name:     "/seeds/world_2.csv",
							Contents: []byte("header_3,header_4\nvalue_5,value_6\nvalue_7,value_8"),
						},
						"/models/world_3.sql": {
							Name:     "/models/world_3.sql",
							Contents: []byte("SELECT * FROM tbd.world_1"),
						},
						"/models/world_4.sql": {
							Name:     "/models/world_4.sql",
							Contents: []byte("SELECT * FROM tbd.world_2 w2 FULL OUTER JOIN tbd.world_3 w3 ON 21.header_3 = w3.header_1"),
						},
					},
				})
			},
			arr: []nodeWithName{
				{"world_1", &servicev1.Seed{
					Name:     "world_1",
					FilePath: "seeds/world_1.csv",
				}},
				{"world_2", &servicev1.Seed{
					Name:     "world_2",
					FilePath: "seeds/world_2.csv",
				}},
				{
					"world_3", &servicev1.Model{
						Name:       "world_3",
						FilePath:   "models/world_3.sql",
						References: []string{"world_1"},
					},
				},
				{
					"world_4", &servicev1.Model{
						Name:       "world_4",
						FilePath:   "models/world_4.sql",
						References: []string{"world_2", "world_3"},
					},
				},
			},
			want:    "WITH\nworld_1 AS (SELECT column1 AS header_1, column2 AS header_2 FROM (VALUES ('value_1', 'value_2'), ('value_3', 'value_4'))),\n world_2 AS (SELECT column1 AS header_3, column2 AS header_4 FROM (VALUES ('value_5', 'value_6'), ('value_7', 'value_8'))),\n world_3 AS (SELECT * FROM world_1) SELECT * FROM (SELECT * FROM world_2 w2 FULL OUTER JOIN world_3 w3 ON 21.header_3 = w3.header_1)",
			wantErr: assert.NoError,
		},
		{
			name:   "array with 4 elements, two seeds, one model, one source",
			config: defaultConfig,
			fsMaker: func() (fs.FS, error) {
				return NewFileSystem(&servicev1.FileSystem{
					Files: map[string]*servicev1.File{
						"/seeds/world_1.csv": {
							Name:     "/seeds/world_1.csv",
							Contents: []byte("header_1,header_2\nvalue_1,value_2\nvalue_3,value_4"),
						},
						"/models/world_3.sql": {
							Name:     "/models/world_3.sql",
							Contents: []byte("SELECT * FROM tbd.world_1"),
						},
						"/models/world_4.sql": {
							Name:     "/models/world_4.sql",
							Contents: []byte("SELECT * FROM tbd.world_2 w2 FULL OUTER JOIN tbd.world_3 w3 ON 21.header_3 = w3.header_1"),
						},
					},
				})
			},
			arr: []nodeWithName{
				{"world_1", &servicev1.Seed{
					Name:     "world_1",
					FilePath: "seeds/world_1.csv",
				}},
				{"world_2", &servicev1.Source{
					Name:     "world_2",
					Path:     "with_schema.world_2",
					FilePath: "models/world_2.yml",
				}},
				{
					"world_3", &servicev1.Model{
						Name:       "world_3",
						FilePath:   "models/world_3.sql",
						References: []string{"world_1"},
					},
				},
				{
					"world_4", &servicev1.Model{
						Name:       "world_4",
						FilePath:   "models/world_4.sql",
						References: []string{"world_2", "world_3"},
					},
				},
			},
			want:    "WITH\nworld_1 AS (SELECT column1 AS header_1, column2 AS header_2 FROM (VALUES ('value_1', 'value_2'), ('value_3', 'value_4'))),\n world_2 AS (SELECT * FROM with_schema.world_2),\n world_3 AS (SELECT * FROM world_1) SELECT * FROM (SELECT * FROM world_2 w2 FULL OUTER JOIN world_3 w3 ON 21.header_3 = w3.header_1)",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsfs, err := tt.fsMaker()
			require.NoError(t, err)

			got, err := convertToSelectStatements(tt.config, fsfs, map[string]*servicev1.Source{}, tt.arr)

			if !tt.wantErr(t, err, fmt.Sprintf("convertToSelectStatements(%v)", tt.arr)) {
				return
			}
			assert.Equalf(t, tt.want, got, "convertToSelectStatements(%v)", tt.arr)
		})
	}
}
