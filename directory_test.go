// scuffold project main.go
package filegenerator

import (
	"testing"
)

func TestDirectory(t *testing.T) {
	d := NewDirectoryGenerator()
	err := d.Add("ABX")
	if err != nil {
		t.Errorf("adding directory failed %v", err)
	} else {
		t.Log("adding directory is possible")
	}
}
