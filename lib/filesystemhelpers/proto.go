package filesystemhelpers

import (
	"io/fs"
	"strings"

	servicev1 "github.com/benfdking/tbd/proto/gen/go/tbd/service/v1"
)

// FSToProtoFileSystem converts an fs.FS to a proto FileSystem so that it can be sent to the frontend.
// TODO should write a test for this
func FSToProtoFileSystem(filesystem fs.FS, root string) (*servicev1.FileSystem, error) {
	out := &servicev1.FileSystem{
		Files: make(map[string]*servicev1.File),
	}
	err := fs.WalkDir(filesystem, root, func(path string, d fs.DirEntry, err error) error {
		destination := strings.TrimPrefix(path, root)
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		file, err := fs.ReadFile(filesystem, path)
		if err != nil {
			return err
		}
		out.Files[destination] = &servicev1.File{
			Name:     destination,
			Contents: file,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}
