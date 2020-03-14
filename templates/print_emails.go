package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type Person struct {
	Name   string
	Emails []string
}

const templ = `The name is {{.Name}}.
{{range .Emails}} An email is {{. | emailExpand}} {{end}}
`

func EmailExpander(args ...interface{}) string {
	ok := false
	var s string

	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}

	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}

	return (substrs[0] + " at " + substrs[1])
}

func main() {
	person := Person{
		Name:   "nam",
		Emails: []string{"namtx.93@gmail.com", "nam@nimblehq.co"},
	}

	t := template.New("Person template")

	t = t.Funcs(template.FuncMap{"emailExpand": EmailExpander})
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
