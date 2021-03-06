package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	url, _ := url.Parse(os.Args[1])

	client := &http.Client{}

	request, _ := http.NewRequest("GET", url.String(), nil)

	request.Header.Add("Accept-Charset", "UTF-8;q=1 ISO-8859-1;q=0")

	response, _ := client.Do(request)

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	chSet := getCharset(response)
	fmt.Printf("Got charset %s\n", chSet)

	if chSet != "UTF-8" {
		fmt.Println("Cannot handle", chSet)
		os.Exit(4)
	}

	var buf [512]byte
	reader := response.Body
	fmt.Println("got body")

	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Println(string(buf[0:n]))
	}

	os.Exit(0)
}

func getCharset(response *http.Response) string {
	contentType := response.Header.Get("Content-Type")

	if contentType == "" {
		return "UTF-8"
	}

	idx := strings.Index(contentType, "charset:")
	if idx == -1 {
		return "UTF-8"
	}

	return strings.Trim(contentType[idx:], " ")
}
