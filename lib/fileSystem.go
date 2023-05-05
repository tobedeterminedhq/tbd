package lib

import (
	"fmt"
	"path/filepath"
	"strings"

	servicev1 "github.com/benfdking/tbd/proto/gen/go/tbd/service/v1"
	"github.com/psanford/memfs"
)

// NewFileSystem generates an in memory filesystem from proto
func NewFileSystem(fs *servicev1.FileSystem) (*memfs.FS, error) {
	rootFS := memfs.New()
	for _, f := range fs.GetFiles() {
		dir, _ := filepath.Split(f.GetName())
		if dir != "/" {
			trimmed := strings.Trim(dir, "/")
			if err := rootFS.MkdirAll(trimmed, 0o755); err != nil {
				return nil, fmt.Errorf("making dir %w", err)
			}
		}
		if err := rootFS.WriteFile(strings.Trim(f.Name, "/"), f.Contents, 0o755); err != nil {
			return nil, fmt.Errorf("writing file %w", err)
		}
	}
	return rootFS, nil
}
