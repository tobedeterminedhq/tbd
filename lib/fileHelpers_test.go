package lib

import (
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	servicev1 "github.com/tobedeterminedhq/tbd/proto_gen/go/tbd/service/v1"
)

func Test_listAllFilesInDirectoryRecursively(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		f       func() (fs.FS, error)
		root    string
		want    []string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "simple example: single file",
			f: func() (fs.FS, error) {
				f := &servicev1.FileSystem{
					Files: map[string]*servicev1.File{
						"seeds/file1.sql": {
							Name:     "seeds/file1.sql",
							Contents: []byte("select 1"),
						},
					},
				}
				return NewFileSystem(f)
			},
			root:    "seeds",
			want:    []string{"seeds/file1.sql"},
			wantErr: assert.NoError,
		},
		{
			name: "simple example: single file, with leading slash",
			f: func() (fs.FS, error) {
				f := &servicev1.FileSystem{
					Files: map[string]*servicev1.File{
						"root/seeds/file1.sql": {
							Name:     "root/seeds/file1.sql",
							Contents: []byte("select 1"),
						},
					},
				}
				return NewFileSystem(f)
			},
			root:    "/root/seeds",
			want:    []string{"root/seeds/file1.sql"},
			wantErr: assert.NoError,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := tt.f()
			require.NoError(t, err)

			got, err := listAllFilesInDirectoryRecursively(f, tt.root)

			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
