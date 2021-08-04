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
		//buf := make([]byte, 1024)
		//_, err = conn.Read(buf)
		reader := week9.NewReader(conn)
		p, err := reader.Decode()
		conn.Close()
		checkErr(err)
		//p := &week9.Protocol{}
		//p.Decode(buf)
		fmt.Println(p)
		fmt.Println(p.Ver)
		fmt.Println(p.Op)
		fmt.Println(p.SeqId)
		fmt.Println(string(p.Body))
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
}
