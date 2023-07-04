package main

import (
"bufio"
"io"
"net"
"os"
"strconv"
)

/*Recovery our input message into a buffer*/
func inputListener() []byte {
	buf := make([]byte, 512)
	readerInput := bufio.NewReader(os.Stdin)
	_, err := readerInput.Read(buf)
	if err != nil {
		panic("Error reading input.")
	}
	return buf
}

func main() {
	if len(os.Args) != 3 {
		println("Usage: ", os.Args[0], " <host> <port>\n")
		os.Exit(0)
	}

	//Recovery the port.
	port, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic("Error during the port recovery\n")
	}
	println(port)

	/*Join the adresse*/
	addr := os.Args[1] + ":" + strconv.Itoa(port)
	println(addr)

	/*  sources -- https://golang.org/pkg/net/  */

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic("Error connecting " + addr)
	}
	defer conn.Close()

	go io.Copy(os.Stdout, conn)

	r := bufio.NewReader(os.Stdin)
	for {
		p, err := r.ReadSlice('\n')
		if err != nil {
			panic("Error reading output.")
		}
		conn.Write(p)
	}
}
