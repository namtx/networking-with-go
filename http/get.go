package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}

	url := os.Args[1]
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	b, _ := httputil.DumpResponse(response, false)
	fmt.Println(string(b))

	contentTypes := response.Header["Content-Type"]

	if !acceptableCharset(contentTypes) {
		fmt.Println(contentTypes)
		os.Exit(4)
	}

	var buf [512]byte
	reader := response.Body

	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Println(string(buf[0:n]))
	}

	os.Exit(0)
}

func acceptableCharset(contentTypes []string) bool {
	for _, contentType := range contentTypes {
		if strings.Index(strings.ToUpper(contentType), "UTF-8") != 1 {
			return false
		}
	}

	return true
}
