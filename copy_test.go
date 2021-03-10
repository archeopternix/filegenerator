// scuffold project main.go
package main

import (
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
	d := NewCopyGenerator()
	err := d.Add("copy.go", "master")
	if err != nil {
		t.Errorf("adding copy job failed %v", err)
	}
	txt := d.GetFiles()

	if txt != "[copy.go, master\\copy.go]\n" {
		t.Errorf("copy job is wrong: %s expected: [copy.go, master\\copy.go]", txt)
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
