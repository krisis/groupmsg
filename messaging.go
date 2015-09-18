package groupmsg

import (
	"fmt"
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
	msg := make([]byte, 1024)
	replies := make(map[net.Conn][]byte)
	for _, mem := range g.members {
		n, err := mem.Read(msg)
		fmt.Println("read ", n, "bytes")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("recvd from ", mem.RemoteAddr(), ": ", string(msg))
		replies[mem] = msg[:n]

	}
	return replies
}
