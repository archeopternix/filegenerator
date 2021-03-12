// scuffold project main.go
package filegenerator

import (
	"io"
	"log"
	"os"
)

func pathExist(name string) bool {
	_, err := os.Lstat(name)
	if err != nil {
		return false
	}
	return true
}

// fileExist returns whether the given file or directory exists
func fileExist(fname string) bool {
	_, err := os.Stat(fname)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// copyFile copies the content from sourcefile to destfile
func copyFile(sourcefile, destfile string) error {
	var source, dest *os.File
	var err error

	// open source file
	source, err = os.Open(sourcefile)
	if err != nil {
		return err
	}
	defer source.Close()

	// overwrite or new file
	exist := fileExist(destfile)

	// create target file
	dest, err = os.Create(destfile)
	if err != nil {
		return err
	}
	defer dest.Close()
	_, err = io.Copy(dest, source)
	if err != nil {
		return err
	}
	if exist {
		log.Printf("file overwritten %s\n", destfile)
	} else {
		log.Printf("file generated %s\n", destfile)
	}
	return nil
}
