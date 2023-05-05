package filesystemhelpers

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// WriteFSToDisk writes a filesystem to disk. It does so by recursively walking the filesystem and writing each file to disk.
//
// root defines the starting point of the walk. It is relative to the filesystem.
// TODO should write a test
func WriteFSToDisk(filesystem fs.FS, root string, dst string) error {
	return fs.WalkDir(filesystem, root, func(path string, d fs.DirEntry, err error) error {
		destination := filepath.Join(dst, strings.TrimPrefix(path, root))
		if err != nil {
			return err
		}
		if d.IsDir() {
			return os.MkdirAll(destination, os.ModePerm)
		}
		return copyFile(filesystem, path, destination)
	})
}

// copyFile writes the file at src in the filesystem to dst on disk.
func copyFile(filesystem fs.FS, file string, dst string) error {
	filePtr, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("error creating file '%s': %w", file, err)
	}
	defer filePtr.Close() // close the file
	toWrite, err := fs.ReadFile(filesystem, file)
	if err != nil {
		return err
	}
	_, err = filePtr.Write(toWrite)
	if err != nil {
		return err
	}
	return nil
}
