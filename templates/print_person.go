package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
	Jobs   []*Job
}

type Job struct {
	Employer string
	Role     string
}

const tmpl = `The name is {{.Name}}. The age is {{.Age}}. 
{{range .Emails}}
        An Email is {{.}}
{{end}}

{{with .Jobs}}
    {{range .}}
        An employer is {{.Employer}}
        and the role is {{.Role}}
    {{end}}
{{end}}`

func main() {
	job1 := Job{Employer: "Nimblehq", Role: "Ruby developer"}

	person := Person{
		Name:   "Tran Xuan Nam",
		Age:    27,
		Emails: []string{"namtx.93@gmail.com", "nam@nimblehq.co"},
		Jobs:   []*Job{&job1},
	}

	t := template.New("Person template")
	t, err := t.Parse(tmpl)
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
