package lib

import (
	"fmt"
	"net"
)

type Handler interface {
	Handle(net.Conn)
}

type EchoHandler struct {
}

func (hndlr *EchoHandler) Handle(conn net.Conn) {
	buf := make([]byte, 1024)
	defer conn.Close()
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error Reading:", err.Error())
		return
	}

	conn.Write([]byte("I received: "))

	conn.Write(buf[:n])
}

func NewHandler() Handler {
	return &EchoHandler{}
}
