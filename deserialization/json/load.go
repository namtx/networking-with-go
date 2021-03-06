package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

type Person struct {
	Name  Name
	Email []Email
}

func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family

	for _, v := range p.Email {
		s += "\n" + v.Kind + ": " + v.Address
	}

	return s
}

func main() {
	var person Person

	loadJSON("person.json", &person)

	fmt.Println("Person", person.String())
}

func loadJSON(filename string, key interface{}) {
	inFile, err := os.Open(filename)
	checkError(err)

	decoder := json.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err)

	inFile.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
}
