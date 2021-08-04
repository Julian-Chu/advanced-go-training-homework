package main

import (
	"fmt"
	"github.com/Julian-Chu/advanced-go-training-homework/week9"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp4", ":9600")
	checkErr(err)
	conn, err := net.DialTCP("tcp", nil, addr)
	checkErr(err)
	defer conn.Close()
	p := week9.NewProto(1, 2, 3, []byte("hello"))
	buf := p.Encode()
	fmt.Println("package len:", len(buf))
	_, err = conn.Write(buf)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		os.Exit(1)
	}
}
