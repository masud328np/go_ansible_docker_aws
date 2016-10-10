package lib

import (
	"fmt"
	"net"
)

type Server interface {
	StartListening(string) bool
}

type NetWrapper interface {
	Listen(string, string) (net.Listener, error)
}

type EchoServer struct {
	NetWrapper     NetWrapper
	RequestHandler Handler
}

func (srvr *EchoServer) StartListening(addrPort string) bool {

	listener, err := srvr.NetWrapper.Listen("tcp", "localhost:"+addrPort)
	if err != nil {
		fmt.Sprintf("error :", err.Error())
		return false
	}

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		//if conn != nil {
		srvr.RequestHandler.Handle(conn)
		//}
		if err != nil || conn == nil {
			break
		}
	}
	return true
}

type EchoListener struct {
}

func (l EchoListener) Listen(proto string, addrPort string) (listener net.Listener, err error) {
	fmt.Printf("Listener")
	listener, err = net.Listen(proto, addrPort)
	return
}

func NewListener() NetWrapper {
	return EchoListener{}
}

func NewServer(handler Handler, netw NetWrapper) Server {
	return &EchoServer{RequestHandler: handler, NetWrapper: netw}
}
