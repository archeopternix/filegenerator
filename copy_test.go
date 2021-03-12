// scuffold project main.go
package filegenerator

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestCopyAdd(t *testing.T) {
	d := NewCopyGenerator()
	err := d.Add("copy.go", "ToDo")
	if err != nil {
		t.Errorf("adding copy job failed %v", err)
	} else {
		t.Log("adding copy job is possible")
	}
}

func TestCopyGetFiles(t *testing.T) {
	const cfile = "copy.go"
	const cdir = "master"
	output := fmt.Sprintf("[%s, %s]", cfile, filepath.Join(cdir, filepath.Base(cfile)))
	t.Log("separator: " + string(os.PathSeparator))

	d := NewCopyGenerator()
	err := d.Add("copy.go", "master")
	if err != nil {
		t.Errorf("adding copy job failed %v", err)
	}
	txt := d.GetFiles()

	if txt != output+"\n" {
		t.Errorf("copy job is wrong: %s expected: %s", txt, output)
	} else {
		t.Log("copy job is right")
	}

}

func TestCopyAddFileNotExist(t *testing.T) {
	d := NewCopyGenerator()
	err := d.Add("abc.tmp", "ToDo")
	if err != nil {
		t.Logf("adding copy job failed %v", err)
	} else {
		t.Error("adding copy job is possible")
	}
}

func TestCopyRun(t *testing.T) {
	d := NewCopyGenerator()
	err := d.Add("copy.go", "ToDo")
	if err != nil {
		t.Errorf("adding copy job failed %v", err)
	} else {
		t.Log("adding copy job is possible")
	}

	err = d.Run()
	if err != nil {
		t.Logf("run copy job failed by design %v", err)

	} else {
		t.Error("run copy job must fail")
	}

}
