package filegenerator

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"
)

const letter = `Dear {{.Name}}, {{if .Attended}}It was a pleasure to see you at the wedding.{{else}}It is a shame you couldn't make it to the wedding.{{end}}{{with .Gift}} Thank you for the lovely {{title .}}.{{end}}`

// Prepare some data to insert into the template.
type Recipient struct {
	Name, Gift string
	Attended   bool
}

var recipient = Recipient{
	"Aunt Mildred", "bone china tea set", false}

func ExampleTemplate() {

	funcMap := template.FuncMap{
		"title": strings.Title,
	}

	tg := NewTemplateGenerator(funcMap)

	if err := tg.Add(tg.Template("letter").Parse(letter)); err != nil {
		log.Fatal(err)
	}

	var b bytes.Buffer

	if tg.ParseWriter("letter", &b, recipient) != nil {
		log.Fatal("parse failed")
	}

	if err := tg.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(b.String())
	// Output: Dear Aunt Mildred, It is a shame you couldn't make it to the wedding. Thank you for the lovely Bone China Tea Set.
}
