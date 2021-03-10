// scuffold project main.go
package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func TestTemplateGenerator() {
	const letter = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
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

	t := NewTemplateGenerator(nil)

	err := t.AddTemplate(template.New("letter").Parse(letter))
	if err != nil {
		log.Fatal(err)
	}

	if t.ParseWriter("letter", os.Stdout, recipient) != nil {
		log.Fatal("parse failed")
	}
	e.AddGenerator(t)
	e.Run()

}

func main() {
	fmt.Println("start")
	TestTemplateGenerator()

	e := new(Engine)
	/*
		d := NewDirectoryGenerator()
		d.Add("generator/misc")

		c := NewCopyGenerator()
		if err := c.Add("doc.go", "generator/misc/"); err != nil {
			log.Fatal(err)
		}

		e.AddGenerator(d)
		e.AddGenerator(c)
	*/

	if err := e.Run(); err != nil {
		log.Fatal(err)
	}

	log.Println("done")
}

/*
func main() {
	//	log.Println(generator.GetAllFilesWithExt("C:\\Go", ".md"))

	e, err := model.App.NewEntity("Sample1", model.Regular)
	if err != nil {
		log.Println(err)
	}

	f1 := model.Field{Name: "ID", FieldType: model.Integer, Required: true}
	e.AddField(&f1)
	f2 := model.Field{Name: "Name", FieldType: model.String}
	e.AddField(&f2)

	_, err = model.App.NewEntity("Sample2", model.Regular)
	if err != nil {
		log.Fatal(err)
	}

	_, err = model.App.NewRelation("Sample1", "Sample2", model.One2many)
	if err != nil {
		log.Fatal(err)
	}

	ds, er := repository.NewYAMLDatastore("abc.yaml")
	if er != nil {
		log.Fatal(er)
	}
	ds.SaveAllData(model.App)


		files := []string{"C:\\go\\robots.txt", "C:\\go\\readme.md"}
		dg := generator.NewCopyGenerator(files, "c:\\Users\\A.Eisner\\go\\src")
		err := dg.Setup()
		if err != nil {
			log.Panic(err)
		}
		err = dg.Run()
		if err != nil {
			log.Panic(err)
		}
		log.Println("Finish")

}
*/
