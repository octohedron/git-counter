package main

import (
	"io/fs"
	"testing"
	"time"
)

var mockDirs = []string{
	"/Users/test",
	"/Users/test",
	"/Users/test",
	"/Users/test",
	"/Users/test",
	"/Users/test",
	"/Users/test",
}

type mockIoHandler struct{}

type fileInfoMock struct{}

func (f fileInfoMock) Name() string       { return "test" }
func (f fileInfoMock) Size() int64        { return int64(100) }
func (f fileInfoMock) Mode() fs.FileMode  { return fs.FileMode(1) }
func (f fileInfoMock) ModTime() time.Time { return time.Now() }
func (f fileInfoMock) IsDir() bool        { return false }
func (f fileInfoMock) Sys() interface{}   { return nil }

func (i mockIoHandler) ReadDir(path string) ([]fs.FileInfo, error) {
	files := make([]fs.FileInfo, 0)
	for i := 0; i < 10; i++ {
		files = append(files, fileInfoMock{})
	}
	return files, nil
}

func (i mockIoHandler) Stat(path string) (fs.FileInfo, error) {
	return fileInfoMock{}, nil
}

func tFatal(err error, msg string, t *testing.T) {
	if err != nil {
		t.Fatal(msg)
	}
}

func TestLoadDirectories(t *testing.T) {
	ioDirHandler := mockIoHandler{}
	projects, err := loadDirectories(ioDirHandler, mockDirs)
	tFatal(err, "Error loading dir", t)
	if len(projects.directories) < 1 {
		t.Fatal("Error, shouldn't be empty")
	}

	type test struct {
		input []string
		want  int
	}

	// Table driven test, each folder has 10 files as implemented in
	// our mock ReadDir() which adds 10 files to each folder
	tests := []test{
		{input: []string{"folder-0"}, want: 10},
		{input: []string{"folder-0", "folder-1"}, want: 20},
		{input: []string{"folder-0", "folder-1", "folder-2"}, want: 30},
	}

	for _, v := range tests {
		ps, err := loadDirectories(ioDirHandler, v.input)
		tFatal(err, "Error loading dir", t)
		if len(ps.directories) != v.want {
			t.Fatal(
				"Error, unexpected value, have",
				len(ps.directories), "want", v.want)
		}
	}
}
