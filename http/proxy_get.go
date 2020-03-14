package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	proxyString := os.Args[1]
	proxyURL, err := url.Parse(proxyString)
	checkError(err)

	rawURL := os.Args[2]
	url, e := url.Parse(rawURL)
	checkError(err)

	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	client := &http.Client{Transport: transport}

	request, err := http.NewRequest("GET", rawURL.String(), nil)
	checkError(err)

	dump, err := httputil.DumpRequest(request, false)
	checkError(err)

	fmt.Println(string(dump))

	response, err := client.Do(request)
	fmt.Println("read ok")

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
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

	if err != nil {

	}

	os.Exit(1)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: ", err.Error())
	}
}
