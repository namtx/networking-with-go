package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name   string
	Emails []string
}

const templ = `{{$name := .Name}}
{{range .Emails}}
    Name is {{$name}}, email is {{.}}
{{end}}`

func main() {
	person := Person{
		Name:   "nam",
		Emails: []string{"namtx.93@gmail.com", "nam@nimblehq.co"},
	}

	t := template.New("Person template")
	t, err := t.Parse(templ)
	checkError(err)

	err = t.Execute(os.Stdout, person)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: ", err.Error())
		os.Exit(1)
	}
}
