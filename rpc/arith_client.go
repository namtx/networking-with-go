package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	serverAddress := os.Args[1]

	client, err := rpc.DialHTTP("tcp", serverAddress)

	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Quotient", args, &quot)
	fmt.Println("Arith: %d*%d=%d\n", args.A, args.B, quot.Quo, quot.Rem)
}
