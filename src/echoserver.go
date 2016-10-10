package main

import "github.com/masud328np/go_ansible_docker_aws/src/lib"

const PORT = "8090"

func main() {

	listener := lib.NewListener()
	handler := lib.NewHandler()
	echoserver := lib.NewServer(handler, listener)
	echoserver.StartListening(PORT)
}
