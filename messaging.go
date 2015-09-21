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

func sendmsg(mem net.Conn, msg []byte, done chan bool) {
	go func() {
		mem.Write(msg)
		done <- true
	}()
}

func (g *Group) SendMsg(msg []byte) {
	done := make(chan bool)
	for _, mem := range g.members {
		sendmsg(mem, msg, done)
	}

	for i := 0; i < len(g.members); i++ {
		<-done
	}
}

func recvmsg(mem net.Conn, replies map[net.Conn][]byte, done chan bool) {
	go func() {
		msg := make([]byte, 1024)
		n, err := mem.Read(msg)
		if err != nil {
			fmt.Println(err)
		}
		replies[mem] = msg[:n]
		done <- true
	}()
}

func (g *Group) RecvMsg() map[net.Conn][]byte {
	replies := make(map[net.Conn][]byte)
	done := make(chan bool)
	for _, mem := range g.members {
		recvmsg(mem, replies, done)
	}

	for i := 0; i < len(g.members); i++ {
		<-done
	}
	return replies
}
