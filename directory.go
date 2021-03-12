// scuffold project main.go
package filegenerator

import (
	"fmt"
	"log"
	"os"
)

// DirectoryGenerator creates new directories in the file system or skips when
// already created. DirectoryGenerator has to be added to Engine
type DirectoryGenerator struct {
	dirs []string // will hold all directories that should be created
}

// NewDirectoryGenerator creates a new instance of Directory Generator
func NewDirectoryGenerator() *DirectoryGenerator {
	return new(DirectoryGenerator)
}

// Add adds a new path to the backlog
func (d *DirectoryGenerator) Add(path string) error {
	d.dirs = append(d.dirs, path)
	return nil
}

//Run executes to creation of all paths added to the DirectoryGenerator
func (d DirectoryGenerator) Run() error {
	for _, path := range d.dirs {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("directory generator: %v", err)
		}
		log.Printf("directory created: '%v'\n", path)
	}
	return nil
}
