package main

import (
	"fmt"
	"github.com/Julian-Chu/advanced-go-training-homework/week9"
	"net"
	"os"
)

func main() {
	port := ":9600"
	listener, err := net.Listen("tcp", port)
	checkErr(err)
	fmt.Printf("Listening on %s\n", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
		reader, close := week9.NewReader(conn)
		p, err := reader.Decode()
		checkErr(err)
		err = close()
		fmt.Println(err)
		fmt.Printf("version: %d\n", p.Ver)
		fmt.Printf("operation: %d\n", p.Op)
		fmt.Printf("sequence ID: %d\n\n", p.SeqId)
		fmt.Printf("message: %s\n\n", string(p.Body))
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
}
