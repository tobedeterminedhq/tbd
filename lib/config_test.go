package lib

import (
	"fmt"
	"io"
	"io/fs"
	"regexp"
	"strings"
	"testing"

	servicev1 "github.com/benfdking/tbd/proto/gen/go/tbd/service/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func Test_parseColumnTests(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		pathMap      PathMap
		customTests  map[string]*servicev1.CustomTest
		projectFiles map[string]*servicev1.ProjectFile
		want         map[string]*servicev1.Test
		wantErr      assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases. Unrecognized test type should return an error.
		// TODO: Add tests for wrong info in relationship
		{
			name:        "simple example for model just references",
			customTests: map[string]*servicev1.CustomTest{},
			pathMap:     PathMap{"example": "example", "users": "users", "users_other": "users_other_path"},
			projectFiles: map[string]*servicev1.ProjectFile{
				"example/example.sql": {
					Models: []*servicev1.ProjectFile_Model{
						{
							Name: "example",
							Columns: []*servicev1.ProjectFile_Column{
								{
									Name:        "id",
									Description: "test",
									Tests: []*servicev1.ProjectFile_Column_ColumnTest{
										{
											Type: "relationship",
											Info: map[string]string{
												"model":  "users",
												"column": "id",
											},
										},
										{
											Type: "relationship",
											Info: map[string]string{
												"model":  "users_other",
												"column": "id",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: map[string]*servicev1.Test{
				"test_example_id_relationship_users_id": {
					TestType: &servicev1.Test_Relationship{
						Relationship: &servicev1.TestRelationship{
							SourceModel:  "example",
							SourcePath:   "example",
							SourceColumn: "id",
							TargetModel:  "users",
							TargetPath:   "users",
							TargetColumn: "id",
							FilePath:     "example/example.sql",
						},
					},
				},
				"test_example_id_relationship_users_other_id": {
					TestType: &servicev1.Test_Relationship{
						Relationship: &servicev1.TestRelationship{
							SourceModel:  "example",
							SourcePath:   "example",
							SourceColumn: "id",
							TargetModel:  "users_other",
							TargetPath:   "users_other_path",
							TargetColumn: "id",
							FilePath:     "example/example.sql",
						},
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "simple example for model",
			customTests: map[string]*servicev1.CustomTest{
				"gte": {
					FilePath: "custom_tests/gte.sql",
					Name:     "gte",
					Sql:      "SELECT * FROM {{ .Model }} WHERE {{ .Column }} < {{ .Value }}",
				},
			},
			pathMap: PathMap{"example": "example", "users": "users", "users_other": "users_other_path"},
			projectFiles: map[string]*servicev1.ProjectFile{
				"example/example.sql": {
					Models: []*servicev1.ProjectFile_Model{
						{
							Name: "example",
							Columns: []*servicev1.ProjectFile_Column{
								{
									Name:        "id",
									Description: "test",
									Tests: []*servicev1.ProjectFile_Column_ColumnTest{
										{
											Type: "not_null",
										},
										{
											Type: "unique",
										},
										{
											Type: "relationship",
											Info: map[string]string{
												"model":  "users",
												"column": "id",
											},
										},
										{
											Type: "relationship",
											Info: map[string]string{
												"model":  "users_other",
												"column": "id",
											},
										},
										{
											Type: "accepted_values",
											Info: map[string]string{
												"values": "1,2,3",
											},
										},
										{
											Type: "gte",
											Info: map[string]string{
												"value": "0",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: map[string]*servicev1.Test{
				"test_example_id_unique": {
					TestType: &servicev1.Test_Unique{
						Unique: &servicev1.TestUnique{
							Model:    "example",
							Path:     "example",
							Column:   "id",
							FilePath: "example/example.sql",
						},
					},
				},
				"test_example_id_not_null": {
					TestType: &servicev1.Test_NotNull{
						NotNull: &servicev1.TestNotNull{
							Model:    "example",
							Path:     "example",
							Column:   "id",
							FilePath: "example/example.sql",
						},
					},
				},
				"test_example_id_relationship_users_id": {
					TestType: &servicev1.Test_Relationship{
						Relationship: &servicev1.TestRelationship{
							SourceModel:  "example",
							SourcePath:   "example",
							SourceColumn: "id",
							TargetModel:  "users",
							TargetPath:   "users",
							TargetColumn: "id",
							FilePath:     "example/example.sql",
						},
					},
				},
				"test_example_id_relationship_users_other_id": {
					TestType: &servicev1.Test_Relationship{
						Relationship: &servicev1.TestRelationship{
							SourceModel:  "example",
							SourcePath:   "example",
							SourceColumn: "id",
							TargetModel:  "users_other",
							TargetPath:   "users_other_path",
							TargetColumn: "id",
							FilePath:     "example/example.sql",
						},
					},
				},
				"test_example_id_accepted_values": {
					TestType: &servicev1.Test_AcceptedValues{
						AcceptedValues: &servicev1.TestAcceptedValues{
							Model:          "example",
							Path:           "example",
							Column:         "id",
							FilePath:       "example/example.sql",
							AcceptedValues: []string{"1", "2", "3"},
						},
					},
				},
				"test_example_id_gte": {
					TestType: &servicev1.Test_CustomColumn{
						CustomColumn: &servicev1.TestCustomColumn{
							TestFilePath: "custom_tests/gte.sql",
							TestName:     "gte",
							OriginalSql:  "SELECT * FROM {{ .Model }} WHERE {{ .Column }} < {{ .Value }}",
							Model:        "example",
							Path:         "example",
							Column:       "id",
							Info: map[string]string{
								"value": "0",
							},
							RenderedSql: "SELECT * FROM example WHERE id < 0",
						},
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name:    "simple example for source",
			pathMap: PathMap{"example": "example", "users": "users", "users_other": "users_other_path"},
			customTests: map[string]*servicev1.CustomTest{
				"gte": {
					FilePath: "custom_tests/gte.sql",
					Name:     "gte",
					Sql:      "SELECT * FROM {{ .Model }} WHERE {{ .Column }} < {{ .Value }}",
				},
			},
			projectFiles: map[string]*servicev1.ProjectFile{
				"example/example.sql": {
					Sources: []*servicev1.ProjectFile_Source{
						{
							Name: "example",
							Path: "example_123",
							Columns: []*servicev1.ProjectFile_Column{
								{
									Name:        "id",
									Description: "test",
									Tests: []*servicev1.ProjectFile_Column_ColumnTest{
										{
											Type: "not_null",
										},
										{
											Type: "unique",
										},
										{
											Type: "relationship",
											Info: map[string]string{
												"model":  "users",
												"column": "id",
											},
										},
										{
											Type: "relationship",
											Info: map[string]string{
												"model":  "users_other",
												"column": "id",
											},
										},
										{
											Type: "accepted_values",
											Info: map[string]string{
												"values": "1,2,3",
											},
										},
										{
											Type: "gte",
											Info: map[string]string{
												"value": "0",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: map[string]*servicev1.Test{
				"test_example_id_unique": {
					TestType: &servicev1.Test_Unique{
						Unique: &servicev1.TestUnique{
							Model:    "example",
							Path:     "example_123",
							Column:   "id",
							FilePath: "example/example.sql",
						},
					},
				},
				"test_example_id_not_null": {
					TestType: &servicev1.Test_NotNull{
						NotNull: &servicev1.TestNotNull{
							Model:    "example",
							Path:     "example_123",
							Column:   "id",
							FilePath: "example/example.sql",
						},
					},
				},
				"test_example_id_relationship_users_id": {
					TestType: &servicev1.Test_Relationship{
						Relationship: &servicev1.TestRelationship{
							SourceModel:  "example",
							SourcePath:   "example_123",
							SourceColumn: "id",
							TargetModel:  "users",
							// TODO Show that this can be a path to a source as well
							TargetPath:   "users",
							TargetColumn: "id",
							FilePath:     "example/example.sql",
						},
					},
				},
				"test_example_id_relationship_users_other_id": {
					TestType: &servicev1.Test_Relationship{
						Relationship: &servicev1.TestRelationship{
							SourceModel:  "example",
							SourcePath:   "example_123",
							SourceColumn: "id",
							TargetModel:  "users_other",
							TargetPath:   "users_other_path",
							TargetColumn: "id",
							FilePath:     "example/example.sql",
						},
					},
				},
				"test_example_id_accepted_values": {
					TestType: &servicev1.Test_AcceptedValues{
						AcceptedValues: &servicev1.TestAcceptedValues{
							Model:          "example",
							Path:           "example_123",
							Column:         "id",
							FilePath:       "example/example.sql",
							AcceptedValues: []string{"1", "2", "3"},
						},
					},
				},
				"test_example_id_gte": {
					TestType: &servicev1.Test_CustomColumn{
						CustomColumn: &servicev1.TestCustomColumn{
							TestFilePath: "custom_tests/gte.sql",
							TestName:     "gte",
							OriginalSql:  "SELECT * FROM {{ .Model }} WHERE {{ .Column }} < {{ .Value }}",
							Model:        "example",
							Path:         "example_123",
							Column:       "id",
							Info: map[string]string{
								"value": "0",
							},
							RenderedSql: "SELECT * FROM example_123 WHERE id < 0",
						},
					},
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseColumnTests(tt.customTests, tt.projectFiles, tt.pathMap)

			if !tt.wantErr(t, err, fmt.Sprintf("parseColumnTests(%v)", tt.projectFiles)) {
				return
			}
			assert.Equalf(t, tt.want, got, "parseColumnTests(%v)", tt.projectFiles)
		})
	}
}

func TestParseConfig(t *testing.T) {
	t.Parallel()

	value := "hello"
	defaultValue := "tbd"

	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *servicev1.Configuration
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "valid config with specified schema",
			args: args{
				r: strings.NewReader(`
schema_name: "hello"
model_paths: ["models"]
seed_paths: ["seeds"]
test_paths: ["tests"]
custom_test_paths: ["custom_tests"]
`),
			},
			want: &servicev1.Configuration{
				SchemaName:      &value,
				ModelPaths:      []string{"models"},
				SeedPaths:       []string{"seeds"},
				TestPaths:       []string{"tests"},
				CustomTestPaths: []string{"custom_tests"},
			},
			wantErr: assert.NoError,
		},
		{
			name: "valid config with no specified schema",
			args: args{
				r: strings.NewReader(`
model_paths: ["models"]
seed_paths: ["seeds"]
test_paths: ["tests"]
custom_test_paths: ["custom_tests"]
`),
			},
			want: &servicev1.Configuration{
				SchemaName:      &defaultValue,
				ModelPaths:      []string{"models"},
				SeedPaths:       []string{"seeds"},
				TestPaths:       []string{"tests"},
				CustomTestPaths: []string{"custom_tests"},
			},
			wantErr: assert.NoError,
		},
		{
			name: "invalid config with bad specified schema",
			args: args{
				r: strings.NewReader(`
schema_name: 1
model_paths: ["models"]
seed_paths: ["seeds"]
test_paths: ["tests"]
custom_test_paths: ["custom_tests"]
`),
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseConfig(tt.args.r)
			if !tt.wantErr(t, err, fmt.Sprintf("ParseConfig(%v)", tt.args.r)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ParseConfig(%v)", tt.args.r)
		})
	}
}

// TODO Test that the replacement works as expected
func Test_returnReferenceSearch(t *testing.T) {
	type args struct {
		schemaName string
	}
	tests := []struct {
		name    string
		args    args
		want    *regexp.Regexp
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := returnReferenceSearch(tt.args.schemaName)
			if !tt.wantErr(t, err, fmt.Sprintf("returnReferenceSearch(%v)", tt.args.schemaName)) {
				return
			}
			assert.Equalf(t, tt.want, got, "returnReferenceSearch(%v)", tt.args.schemaName)
		})
	}
}

func Test_parseModel(t *testing.T) {
	tests := []struct {
		name    string
		fsfs    fs.FS
		c       *servicev1.Configuration
		ms      ModelDefinitions
		path    string
		want    *servicev1.Model
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "valid model",
			// TODO Should refactor this so that over time it's easier to build fresh model
			ms: ModelDefinitions{
				"example": {
					Name:        "example",
					Description: "This is an example model",
				},
			},
			fsfs: func() fs.FS {
				f, err := NewFileSystem(&servicev1.FileSystem{Files: map[string]*servicev1.File{
					"/models/example.sql": {
						Name: "/models/example.sql",
						Contents: []byte(`
WITH shifts AS (SELECT employee_id,
                       shift_date,
                       shift
                FROM tbd.stg_shifts
                ),
     shift_details AS (SELECT shift AS shift_name,
                              start_time,
                              end_time
                       FROM tbd.shift_hours
                       )

SELECT s.employee_id AS employee_id,
       s.shift AS shift,
       datetime(s.shift_date, sd.start_time) AS shift_start,
       datetime(s.shift_date, sd.end_time)   AS shift_end
FROM shifts s
         INNER JOIN shift_details sd
                    ON s.shift = sd.shift_name
`),
					},
					"/models/example.yaml": {
						Name: "/models/example.yaml",
						Contents: []byte(`
models:
  - name: example
    description: description of the model
`),
					},
				}})
				require.NoError(t, err)
				return f
			}(),
			c: &servicev1.Configuration{
				SchemaName: proto.String("tbd"),
			},
			path:    "models/example.sql",
			wantErr: assert.NoError,
			want: &servicev1.Model{
				Name:        "example",
				Description: "This is an example model",
				References:  []string{"stg_shifts", "shift_hours"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseModel(tt.fsfs, tt.ms, tt.c, tt.path)
			if !tt.wantErr(t, err, fmt.Sprintf("parseModel(%v, %v, %v)", tt.fsfs, tt.c, tt.path)) {
				return
			}

			assert.Equalf(t, tt.want.GetName(), got.GetName(), "parseModel(%v, %v, %v)", tt.fsfs, tt.c, tt.path)
			assert.Equalf(t, tt.want.GetReferences(), got.GetReferences(), "parseModel(%v, %v, %v)", tt.fsfs, tt.c, tt.path)
		})
	}
}
