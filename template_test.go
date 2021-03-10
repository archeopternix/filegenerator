// scuffold project main.go
package main

import (
	"os"
	"testing"
	"text/template"
)

func TestTemplateGenerator(t *testing.T) {
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

	e := new(Engine)

	tg := NewTemplateGenerator(nil)

	err := tg.AddTemplate(template.New("letter").Parse(letter))
	if err != nil {
		t.Fatal(err)
	}

	if tg.ParseWriter("letter", os.Stdout, recipient) != nil {
		t.Fatal("parse failed")
	}
	e.AddGenerator(tg)
	err = e.Run()
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("letter test")
	}

}
