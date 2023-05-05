package lib

import (
	"fmt"
	"io/fs"
	"strings"
)

// listAllFilesInDirectoryRecursively lists all files in a recursive way under a specified root. It works for both 'seeds' or '/seeds'.
func listAllFilesInDirectoryRecursively(f fs.FS, root string) ([]string, error) {
	var files []string
	err := fs.WalkDir(f, strings.TrimPrefix(root, "/"), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error in path %s: %w", path, err)
		}
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walking dir: %w", err)
	}
	return files, nil
}

func byType(t fileTypes) func(s string, _ int) bool {
	return func(s string, _ int) bool {
		return strings.HasSuffix(s, string(t))
	}
}

type fileTypes string

const (
	sqlFileType fileTypes = ".sql"
	csvFileType fileTypes = ".csv"
	yml         fileTypes = ".yml"
)
