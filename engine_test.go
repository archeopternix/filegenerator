// scuffold project main.go
package filegenerator

import (
	"fmt"
	"testing"
)

type TestGenerator struct {
}

func (t TestGenerator) Run() error {
	return fmt.Errorf("test error")
}

func TestRun(t *testing.T) {
	e := NewEngine("Test")

	if err := e.Run(); err != nil {
		t.Fatalf("engine run %v", err)
	} else {
		t.Log("empty generator")
	}
}

func TestGeneratorRunFail(t *testing.T) {
	var tg TestGenerator

	e := NewEngine("Test")
	e.AddGenerator(tg)
	if err := e.Run(); err == nil {
		t.Errorf("engine run with test generator %v", err)
	} else {
		t.Log("error raised within test generator as expected")
	}
}
