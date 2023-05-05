package lib_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/benfdking/tbd/go/lib"
	servicev1 "github.com/benfdking/tbd/proto/gen/go/tbd/service/v1"
	"github.com/stretchr/testify/assert"
)

func TestParseProjectFile(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		bs      []byte
		want    *servicev1.ProjectFile
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "invalid yaml",
			bs:      []byte(`:123678`),
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "simple example",
			bs: []byte(`
models: 
- name: model test
  description: test description for model
  columns: 
    - name: column test
      description: test description for column

sources:
- name: source_test
  description: test description for source
  path: source_test.source_test
  columns:
    - name: column test
      description: test description for sources column
`),
			want: &servicev1.ProjectFile{
				Models: []*servicev1.ProjectFile_Model{
					{
						Name:        "model test",
						Description: "test description for model",
						Columns: []*servicev1.ProjectFile_Column{
							{
								Name:        "column test",
								Description: "test description for column",
							},
						},
					},
				},
				Sources: []*servicev1.ProjectFile_Source{
					{
						Name:        "source_test",
						Description: "test description for source",
						Path:        "source_test.source_test",
						Columns: []*servicev1.ProjectFile_Column{
							{
								Name:        "column test",
								Description: "test description for sources column",
							},
						},
					},
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := bytes.NewReader(tt.bs)

			got, err := lib.ParseProjectFile(r)
			if !tt.wantErr(t, err, fmt.Sprintf("ParseProjectFile(%v)", tt.bs)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ParseProjectFile(%v)", tt.bs)
		})
	}
}
