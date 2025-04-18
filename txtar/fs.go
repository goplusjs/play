package txtar

import (
	"io/fs"
	"path"
	"path/filepath"

	"golang.org/x/tools/txtar"
)

// FileSystem represents a file system.
type iFileSystem interface {
	ReadDir(dirname string) ([]fs.DirEntry, error)
	ReadFile(filename string) ([]byte, error)
	Join(elem ...string) string

	// Base returns the last element of path.
	// Trailing path separators are removed before extracting the last element.
	// If the path is empty, Base returns ".".
	// If the path consists entirely of separators, Base returns a single separator.
	Base(filename string) string

	// Abs returns an absolute representation of path.
	Abs(path string) (string, error)
}

func FS(fs *FileSet, filter func(file string) bool) (*FileSystem, error) {
	a := new(txtar.Archive)
	for _, f := range fs.Files {
		if filter != nil && !filter(f) {
			continue
		}
		a.Files = append(a.Files, txtar.File{Name: f, Data: fs.M[f]})
	}
	fsys, err := txtar.FS(a)
	if err != nil {
		return nil, err
	}
	return &FileSystem{fs, fsys}, nil
}

type FileSystem struct {
	Fset *FileSet
	Sys  fs.FS
}

func (f *FileSystem) ReadDir(dirname string) ([]fs.DirEntry, error) {
	dirs, err := fs.ReadDir(f.Sys, dirname)
	return dirs, err
}

func (f *FileSystem) ReadFile(filename string) ([]byte, error) {
	return fs.ReadFile(f.Sys, filename)
}

func (f *FileSystem) Join(elem ...string) string {
	return path.Join(elem...)
}

func (f *FileSystem) Base(filename string) string {
	return path.Base(filename)
}

func (f *FileSystem) Abs(path string) (string, error) {
	return filepath.Abs(path)
}
