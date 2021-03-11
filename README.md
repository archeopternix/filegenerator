# filegenerator - a toolbox for file generation
[![codecov](https://codecov.io/gh/archeopternix/filegenerator/branch/main/graph/badge.svg?token=NK2N53V1X8)](https://codecov.io/gh/archeopternix/filegenerator)
This toolbox will help you when working with templates and setup of file structures in the context of application generation etc.

## Engine
The Engine will act as a container for different kind of [Generator](##Generator) that have to implement the Interface with the only method:

`Run() error`

Sample code how to setup and run an Engine:

`var e Engine`
`err := e.Run()` 

## Generator
Generator is an interface and has to implemented for concrete use cases
```
 // Generator interface for Engine
type Generator interface {
Run() error // runs the generator
}
```

Implemented Generators are added to the Engine using the method:

`e.AddGenerator([Generator](##Generator))`

The execution of the method Run() executes all added generators

`e.Run()`

### CopyGenerator
CopyGenerator create a new Generator for copying files from one place into a directory specified by the method:

`Add(file, todir string) error` 

### DirectoryGenerator
Creates all directories that are added to the DirectoryGenerator by the method:

`Add(path string) error`

### TemplateGenerator 
Creates an new Generator which uses an optional FuncMap to enrich the available functions that can be directly used within templates

New Template generator will be created using the method:

`func NewTemplateGenerator(fmap template.FuncMap) *TemplateGenerator`

Template creates a *template.Template including a FuncMap when added by calling the NewTemplateGenerator method

`Template(name string) (tpl *template.Template)`

The method Add adds a template to the generator. It takes a template and an error to be used convinently with existing template functions

`Add(tpl *template.Template, err error) error`

Before a template can be executed it has to be parsed. For this purpose there are 2 ways implemented

**ParseWriter** adds the relevant information for getting the right template writing it to an io.Writer. The data is used for template generation

`ParseWriter(name string, wr io.Writer, data interface{}) error`

**ParseFilename** holds the relevant information for getting the right template writing it to a file. The Data is used for template generation

`ParseFilename(name string, file string, data interface{}) error `

Sample:
```
const letter =` 
    Dear {{.Name}},
    {{- if .Attended}}
    It was a pleasure to see you at the wedding.
    {{- else}}
    It is a shame you couldn't make it to the wedding.
    {{- end}}
    {{with .Gift -}}
    {{title .}}
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

func main() {
	funcMap := template.FuncMap{
		"title": strings.Title,
	}

	tg := NewTemplateGenerator(funcMap)

	if err := tg.Add(tg.Template("letter").Parse(letter)); err != nil {
		log.Fatal(err)
	}

	if tg.ParseWriter("letter", os.Stdout, recipient) != nil {
		log.Fatal("parse failed")
	}`

	if err := tg.Run(); err != nil {
		log.Fatal(err)
	}
}
```