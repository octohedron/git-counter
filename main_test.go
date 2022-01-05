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

func TestLoadDirectories(t *testing.T) {
	ioDirHandler := mockIoHandler{}
	projects, err := loadDirectories(ioDirHandler, mockDirs)
	if err != nil {
		t.Fatal("Error loading dir")
	}
	if len(projects.directories) < 1 {
		t.Fatal("Error, should be empty")
	}
}
