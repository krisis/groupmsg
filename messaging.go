package groupmsg

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

type Group struct {
	members []net.Conn
}

func (g *Group) Members() []net.Conn {
	return g.members
}

func (g *Group) String() string {
	return fmt.Sprintf("%#v", g.members)
}

func NewGroup() *Group {
	return &Group{}
}

func (g *Group) AddMember(members ...net.Conn) {
	g.members = append(g.members, members...)
}

func (g *Group) SendMsg(msg []byte) {
	for _, mem := range g.members {
		mem.Write(msg)
	}
}

func (g *Group) RecvMsg() map[net.Conn][]byte {
	var msg bytes.Buffer
	replies := make(map[net.Conn][]byte)
	for _, mem := range g.members {
		n, err := io.Copy(&msg, mem)
		fmt.Println("read ", n, "bytes")
		if err != nil {
			fmt.Printf("failed to copy from buffer\n")
		}
		fmt.Println("recvd from ", mem.RemoteAddr(), ": ", msg.Bytes())
		replies[mem] = msg.Bytes()
		msg.Reset()
	}
	return replies
}
