// scuffold project main.go
package main

import (
	"io/ioutil"
	"strings"
	"testing"
	"text/template"
)

const letter = `
Dear {{.Name}},
{{- if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{- end}}
Best wishes,
Josie
`

// Prepare some data to insert into the template.
type Recipient struct {
	Name, Gift string
	Attended   bool
}

var recipient = Recipient{
	"Aunt Mildred", "bone china tea set", false}

func TestTemplateGenerator(t *testing.T) {
	e := new(Engine)

	tg := NewTemplateGenerator(nil)

	if err := tg.Add(tg.Template("letter").Parse(letter)); err != nil {
		t.Fatal(err)
	}

	if tg.ParseWriter("letter", ioutil.Discard, recipient) != nil {
		t.Fatal("parse failed")
	}

	e.AddGenerator(tg)
	if err := e.Run(); err != nil {
		t.Fatal(err)
	} else {
		t.Log("standard letter ok")
	}
}

func TestTemplateGeneratorParseFilename(t *testing.T) {
	tg := NewTemplateGenerator(nil)

	if err := tg.Add(tg.Template("letter").Parse(letter)); err != nil {
		t.Fatal(err)
	}

	if tg.ParseFilename("letter", "testletter.txt", recipient) != nil {
		t.Fatal("parse failed")
	}

	t.Log("standard letter file parser ok")
}

func TestTemplateGeneratorWrongTemplate(t *testing.T) {
	tg := NewTemplateGenerator(nil)

	let := letter + "{"
	if err := tg.Add(tg.Template("letter").Parse(let)); err != nil {
		t.Fatal("expect to fail due to error in template")
	} else {
		t.Log("expected fail due to error in template")
	}
}

func TestTemplateGeneratorWrongTemplateName(t *testing.T) {
	tg := NewTemplateGenerator(nil)

	err := tg.Add(tg.Template("letter").Parse(letter))
	if err != nil {
		t.Fatal("expect to fail due to error in template")
	}

	if tg.ParseWriter("sheet", ioutil.Discard, recipient) != nil {
		t.Log("expected fail due to wrong template name")
	} else {
		t.Fatal("expect to fail due to wrong template name")
	}
	if err := tg.Run(); err != nil {
		t.Fatal(err)
	} else {
		t.Log("letter ok")
	}
}

func TestTemplateGeneratorFunction(t *testing.T) {
	funcMap := template.FuncMap{
		"title": strings.Title,
	}

	tg := NewTemplateGenerator(funcMap)

	let := letter + "{{title .Name}}"

	if err := tg.Add(tg.Template("letter").Parse(let)); err == nil {
		t.Log("funcmap added and executed")
	} else {
		t.Fatalf("funcmap failed to execute %v", err)
	}

	if err := tg.ParseWriter("letter", ioutil.Discard, recipient); err != nil {
		t.Fatalf("error in parsing template name %v", err)
	} else {
		t.Log("letter incl FuncMap parsed")
	}

	if err := tg.Run(); err != nil {
		t.Fatal(err)
	} else {
		t.Log("letter ok")
	}
}
