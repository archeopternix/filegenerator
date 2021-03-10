// scuffold project main.go
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"text/template"
)

// TemplateGenerator produces content based on go templates
type TemplateGenerator struct {
	templates map[string]template.Template
	funcmap   template.FuncMap
	Data      interface{}
	output    []templateOutput
}

// TemplateOutput holds the relevant information for getting the right template
// writing it to an io.Writer or into a file. The Data is used for template generation
type templateOutput struct {
	Name     string    // Name of the template
	FileName string    // FileName or io.Writer has to be provided
	Writer   io.Writer // write the output
	Data     interface{}
}

// NewTemplateGenerator creates an new Generator which uses an option FuncMap to enrich
// the available functions that can be directly used within templates
func NewTemplateGenerator(fmap *template.FuncMap) *TemplateGenerator {
	tg := new(TemplateGenerator)
	if fmap != nil {
		tg.funcmap = *fmap
	}
	tg.templates = make(map[string]template.Template)
	return tg
}

// AddTemplate adds a template to the generator. It takes a template and an error
// to be used convinently with existing template functions
//    e.g AddTemplate(template.New("letter").Parse(letter))
func (tg *TemplateGenerator) AddTemplate(tpl *template.Template, err error) error {
	var tmpl *template.Template

	if err != nil {
		return err
	}

	name := tpl.Name()
	if tg.funcmap != nil {
		tmpl = tpl.New(name).Funcs(tg.funcmap)

	} else {
		tmpl = tpl.New(name)
	}
	if err != nil {
		return fmt.Errorf("Error in adding templates: %s", err)
	}
	tg.templates[name] = *tmpl
	return nil
}

// ParseWriter adds the relevant information for getting the right template
// writing it to an io.Writer. The data is used for template generation
func (tg *TemplateGenerator) ParseWriter(name string, wr io.Writer, data interface{}) error {
	to := templateOutput{
		Name:   name, // Template name
		Writer: wr,
		Data:   data,
	}
	_, ok := tg.templates[name]
	if !ok {
		return fmt.Errorf("template name not found: '%v'", name)
	}
	tg.output = append(tg.output, to)
	return nil
}

// ParseFilename holds the relevant information for getting the right template
// writing it to a file. The Data is used for template generation
func (tg *TemplateGenerator) ParseFilename(name string, file string, data interface{}) error {
	to := templateOutput{
		Name:     name, // Template name
		FileName: file,
		Writer:   nil,
		Data:     data,
	}
	_, ok := tg.templates[name]
	if !ok {
		return fmt.Errorf("template name not found: '%v'", name)
	}
	tg.output = append(tg.output, to)
	return nil
}

// Run will be called by the 'Engine' and creates the files using the data and
// templates provided
func (tg TemplateGenerator) Run() error {
	for _, to := range tg.output {
		tmpl := tg.templates[to.Name]
		// when io.Writer is provided
		if to.Writer != nil {
			err := tmpl.ExecuteTemplate(to.Writer, to.Name, to.Data)
			if err != nil {
				return fmt.Errorf("templategenerator %v", err)
			}
		} else {
			// create target file
			writer, err := os.Create(to.FileName)
			if err != nil {
				return fmt.Errorf("templategenerator %v", err)
			}
			defer writer.Close()
			err = tmpl.ExecuteTemplate(writer, to.Name, to.Data)
			if err != nil {
				return fmt.Errorf("templategenerator %v", err)
			}

			log.Printf("template '%v' generated\n", to.Name)
		}
	}
	return nil
}
