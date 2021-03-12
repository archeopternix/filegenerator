// scuffold project main.go
package filegenerator

import (
	"testing"
)

func TestPathExist(t *testing.T) {
	if ok := pathExist("."); !ok {
		t.Error("path not found '.'")
	}
	if ok := pathExist("not_a_real_path"); ok {
		t.Error("not expected to find path 'not_a_real_path'")
	}
}

func TestFileExist(t *testing.T) {
	if ok := fileExist("tools.go"); !ok {
		t.Error("path not found '.'")
	}
	if ok := fileExist("not_a_real_file.go"); ok {
		t.Error("not expected to find path 'not_a_real_file.go'")
	}
}

/* throws errors on Unix
func TestCopyFiles(t *testing.T) {
	if err := copyFile("tools.go", "not_a_real_directory\\tools.go"); err == nil {
		t.Error("should not copy to 'not_a_real_directory'")
	}
	if err := copyFile("not_a_real_file.go", "not_a_real_directory\\tools.go"); err == nil {
		t.Error("should not find 'not_a_real_file.go'")
	}
}
*/
