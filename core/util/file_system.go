// Copyright 2025 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package util

import (
	"os"
)

// FileSystem struct
type FileSystem struct {
}

// NewFileSystem creates a new instance
func NewFileSystem() *FileSystem {
	return &FileSystem{}
}

// FileExists reports whether the named file exists
func (fs *FileSystem) FileExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsRegular() {
			return true
		}
	}

	return false
}

// DirExists reports whether the dir exists
func (fs *FileSystem) DirExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsDir() {
			return true
		}
	}

	return false
}

// EnsureDir ensures that directory exists
func (fs *FileSystem) EnsureDir(dirName string, mode int) error {
	err := os.MkdirAll(dirName, os.FileMode(mode))

	if err == nil || os.IsExist(err) {
		return nil
	}

	return err
}
