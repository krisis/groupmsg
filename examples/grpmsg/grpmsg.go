package main

import (
	"bytes"
	"fmt"
	"net"
	"os"

	"github.com/krisis/groupmsg"
)

func client(addrs ...string) {
	g := groupmsg.NewGroup()
	for _, addr := range addrs {
		server, err := net.Dial("tcp", fmt.Sprintf("%s%s", "localhost", addr))
		if err != nil {
			fmt.Println("failed to connect to ", addr, " : ", err)
			fmt.Printf("failed to connect to %s\n", addr)
			break
		}
		g.AddMember(server)
	}

	g.SendMsg([]byte("Hello"))
	msgs := g.RecvMsg()

	for k, v := range msgs {
		fmt.Println(k, ": ", string(v))
	}
}

func handle(conn net.Conn) {
	buf := make([]byte, 1024)
	var msg []byte

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("error while reading from %s\n", conn.RemoteAddr())
	}
	msg = buf[:n]

	if bytes.Equal([]byte("Hello"), msg) {
		_, err := conn.Write([]byte("World"))
		if err != nil {
			fmt.Println(err)
		}

	} else {
		fmt.Println("Unknown message: ", string(msg))
	}
}

func server(addr string) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("failed to listen on %s\n", addr)
		os.Exit(1)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("failed to complete connection\n")
			break
		}

		go handle(conn)

	}
}

func main() {

	if len(os.Args) > 2 && os.Args[1] == "server" {
		server(os.Args[2])
		os.Exit(0)
	}
	if len(os.Args) > 3 && os.Args[1] == "client" {
		client(os.Args[2:]...)
		os.Exit(0)
	}
	fmt.Fprintf(os.Stderr, "Usage: messaging-sample server|client <ADDR> <ADDR> ...\n")
	os.Exit(1)
}
