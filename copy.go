// scuffold project main.go
package main

import (
	"fmt"
	"path/filepath"
)

// FromTo is a helper struct to hold source and target file
type fromTo struct {
	From string
	To   string
}

// CopyGenerator holds all files that should be copied over into a new location
type CopyGenerator struct {
	files []fromTo
}

// NewCopyGenerator createas a new Generator for copying files from one place to another
func NewCopyGenerator() *CopyGenerator {
	return new(CopyGenerator)
}

// Add is used to add another 'file' and target directory 'todir' where the
// file will be copied into
func (d *CopyGenerator) Add(file, todir string) error {
	if !fileExist(file) {
		return fmt.Errorf("copy source '%v' does not exist", file)
	}

	to := filepath.Dir(todir) + "\\" + filepath.Base(file)

	ft := fromTo{From: file, To: to}
	d.files = append(d.files, ft)
	return nil
}

//Run executes to creation of all paths added to the CopyGenerator
func (d CopyGenerator) Run() error {
	for _, ft := range d.files {
		err := copyFile(ft.From, ft.To)
		if err != nil {
			return fmt.Errorf("copygenerator %v", err)
		}
	}
	return nil
}
