package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

var (
	imgId       string
	containerId string
)

func TestMain(m *testing.M) {
	//check if path var is set to set env

	isEnvSet, err := strconv.ParseBool(os.Getenv("envset"))

	if err != nil {
		isEnvSet = false
	}

	if !isEnvSet {
		BuildApp()
		StartContainer()
	}
	fmt.Println("Runnning  UAT....")
	v := m.Run()
	if !isEnvSet {
		StopContainer(strings.TrimSpace(containerId), strings.TrimSpace(imgId))
	}
	os.Exit(v)
}

func TestEcho(t *testing.T) {
	fmt.Print("\t TestEcho - ")

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

/* private function*/
func BuildApp() {
	fmt.Println("0. Building app")

	args := []string{"build", "-a", "cgo", "-o", "../../../docker/echoserver", "../../echoserver.go"}
	cmd := exec.Command("go", args...)
	os.Setenv("GOOS", "linux")
	os.Setenv("CGO_ENABLED", "0")
	cmd.Env = os.Environ()
	_, err := cmd.Output()

	if err != nil {
		panic("hehehe")
		fmt.Println("error :" + err.Error())
		return
	}
}

func StartContainer() {

	imageId, err := BuildContainer()
	imgId = strings.Split(strings.TrimSpace(imageId), ":")[1]

	if err != nil {
		os.Exit(1)
	}

	containerId, err = RunContainer(imgId)

	if err != nil {
		os.Exit(1)
	}
}

func BuildContainer() (imageId string, err error) {
	fmt.Println("1. Building Container")
	result, err := ExecCommand("docker", []string{"build", "-t", "tmp:latest", "-f", "../../../docker/Dockerfile", "-q", "../../../docker/"})
	imageId = string(result)
	return
}

func RunContainer(imgId string) (contId string, err error) {

	fmt.Println("2. Running Container:", containerId)
	result, err := ExecCommand("docker", []string{"run", "-d", "-p", "8090:8090", imgId})
	contId = string(result)
	return
}

func ExecCommand(command string, cmdArgs []string) (output []byte, err error) {
	output, err = exec.Command(command, cmdArgs...).Output()
	return
}

func StopContainer(contId string, imageId string) {
	fmt.Println("3. Stopping Container")

	if _, err := ExecCommand("docker", []string{"stop", contId}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("4. removing Container")

	if _, err := ExecCommand("docker", []string{"rm", "-f", contId}); err != nil {
		os.Exit(1)
	}
	fmt.Println("5. Removing images")
	if _, err := ExecCommand("docker", []string{"rmi", imageId, "-f"}); err != nil {
		os.Exit(1)
	}

}
