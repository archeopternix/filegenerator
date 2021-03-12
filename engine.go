// scuffold project main.go
package filegenerator

import (
	"fmt"
	"log"
)

// Generator interface for Engine
type Generator interface {
	Run() error // runs the generator
}

// Engine holds all Generators and triggers the run. A name is needed to
// identify each individual run
type Engine struct {
	Name      string
	generator []Generator
}

func NewEngine(name string) *Engine {
	e := new(Engine)
	e.Name = name
	return e
}

// AddGenerator adds a new Generator
func (e *Engine) AddGenerator(g Generator) error {
	if len(e.Name) < 1 {
		return fmt.Errorf("Engine needs an name")
	}
	e.generator = append(e.generator, g)

	return nil
}

// Run executes all Run() functions of each generator
func (e *Engine) Run() error {
	if len(e.Name) < 1 {
		return fmt.Errorf("Engine needs a name")
	}

	for i, g := range e.generator {
		if err := g.Run(); err != nil {
			return fmt.Errorf("run %v", err)
		}
	}
	log.Printf("Engine '%s'\n", e.Name)
	return nil
}
