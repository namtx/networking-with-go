package main

import (
	"os"
	"encoding/gob"
	"fmt"
)


type Person struct {
	Name Name
	Email [] Email
}

type Name struct {
	Family string
	Personal string
}

type Email struct {
	Kind string
	Address string
}

func main() {
	person := Person{
		Name: Name{
			Family: "Tran",
			Personal: "Nam",
		},
		Email: []Email{
			Email{
				Kind: "work",
				Address: "nam@nimblehq.co",
			},
		},
	}

	saveGob("person.gob", person)
}

func saveGob(filename string, key interface{}) {
	outFile, err := os.Create(filename)
	checkError(err)

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)

	outFile.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
}
