package uat

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"testing"
)

var isEnvSet bool

func Teardown() {

	if !isEnvSet {
		StopContainer(strings.TrimSpace(containerId), strings.TrimSpace(imgId))
	}
}

func Setup() {
	//check if path var is set to set env

	//if !isEnvSet {
	BuildApp()
	StartContainer()
	//}
}

func TestEcho(t *testing.T) {

	isEnvSet, err := strconv.ParseBool(os.Getenv("envset"))

	if err != nil {
		isEnvSet = false
	}

	if !isEnvSet {
		Setup()
	}

	defer func() {
		if !isEnvSet {
			Teardown()
		}
	}()

	ipaddress := os.Getenv("echoAddress")
	if ipaddress == "" {
		ipaddress = "localhost"
	}

	raddr, err := net.ResolveTCPAddr("tcp", ipaddress+":8090")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	defer conn.Close()
	strEcho := "HI\n"
	_, err = conn.Write([]byte(strEcho))
	if err != nil {
		t.Fail()
	}

	message, _ := bufio.NewReader(conn).ReadString('\n')
	if !strings.Contains(message, strEcho) {
		t.Fail()
	}

}
