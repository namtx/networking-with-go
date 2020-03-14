package main

import (
	"fmt"
	"net"
	"os"
)

/*
	How to send ping message to a hosts
	Ping uses the "echo" command from ICMP protocol
	This is byte-oriented protocol
*/

/*
	Format:
		- the first byte is 8, standing for echo message
		- the second byte is zero
		- The third and fourth byte bytes a checksum on entire messsage
		- The fifth and sixth bytes are arbitrary indentifier
		- The seventh and eight bytes are arbitrary sequence
		- The rest of package is user data
*/

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host", os.Args[0])
		os.Exit(1)
	}

	addr, err := net.ResolveIPAddr("ip", os.Args[1])
	if err != nil {
		fmt.Println("Resolution error: ", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialIP("ip4:icmp", addr, addr)
	checkError(err)

	var msg [512]byte
	msg[0] = 8 // echo

	msg[1] = 0

	msg[2] = 0
	msg[3] = 0

	msg[4] = 0
	msg[5] = 13

	msg[6] = 0
	msg[7] = 37

	len := 8

	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)
	_, err = conn.Write(msg[0:len])
	checkError(err)
	_, err = conn.Read(msg[0:])
	checkError(err)

	fmt.Println("got response")

	if msg[5] == 13 {
		fmt.Println("indentifier matches")
	}

	if msg[7] == 37 {
		fmt.Println("sequence matches")
	}

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
}

func checkSum(msg []byte) uint16 {
	sum := 0

	for n := 1; n < len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}

	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)

	var answer uint16 = uint16(^sum)

	return answer
}
