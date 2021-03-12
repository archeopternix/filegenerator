// scuffold project main.go
package filegenerator

import (
	"fmt"
)

// Generator interface for Engine
type Generator interface {
	Run() error // runs the generator
}

// Engine holds all Generators and triggers the run
type Engine struct {
	generator []Generator
}

// AddGenerator adds a new Generator
func (e *Engine) AddGenerator(g Generator) error {
	e.generator = append(e.generator, g)

	return nil
}

// Run executes all Run() functions of each generator
func (e *Engine) Run() error {
	for _, g := range e.generator {
		if err := g.Run(); err != nil {
			return fmt.Errorf("run %v", err)
		}
	}
	return nil
}
