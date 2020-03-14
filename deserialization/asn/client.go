package main

import (
	"encoding/asn1"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

func main() {
	service := ":1200"

	conn, err := net.Dial("tcp", service)
	checkError(err)

	defer conn.Close()

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	var newtime time.Time

	_, err1 := asn1.Unmarshal(result, &newtime)
	checkError(err1)

	fmt.Println("After marshal/unmarshal: ", newtime.String())

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
}
