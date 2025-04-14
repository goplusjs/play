package txtar

import (
	"io/fs"
	"path"
	"path/filepath"

	"golang.org/x/tools/txtar"
)

// FileSystem represents a file system.
type FileSystem interface {
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

func FS(fs *FileSet) (FileSystem, error) {
	a := new(txtar.Archive)
	for _, f := range fs.Files {
		if f == "go.mod" || f == "gop.mod" {
			continue
		}
		a.Files = append(a.Files, txtar.File{Name: f, Data: fs.M[f]})
	}
	fsys, err := txtar.FS(a)
	if err != nil {
		return nil, err
	}
	return &fileSystem{fs, fsys}, nil
}

type fileSystem struct {
	fs   *FileSet
	fsys fs.FS
}

func (f *fileSystem) ReadDir(dirname string) ([]fs.DirEntry, error) {
	dirs, err := fs.ReadDir(f.fsys, dirname)
	return dirs, err
}

func (f *fileSystem) ReadFile(filename string) ([]byte, error) {
	return fs.ReadFile(f.fsys, filename)
}

func (f *fileSystem) Join(elem ...string) string {
	return path.Join(elem...)
}

func (f *fileSystem) Base(filename string) string {
	return path.Base(filename)
}

func (f *fileSystem) Abs(path string) (string, error) {
	return filepath.Abs(path)
}
