package uat

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	imgId       string
	containerId string
)

/* private function*/
func BuildApp() {
	fmt.Println("0. Building app")

	args := []string{"build", "-a", "-installsuffix", "cgo", "-o", "../../../docker/echoserver", "../../echoserver.go"}
	cmd := exec.Command("go", args...)
	os.Setenv("GOOS", "linux")
	os.Setenv("CGO_ENABLED", "0")
	cmd.Env = os.Environ()
	_, err := cmd.Output()

	if err != nil {
		//panic("hehehe")
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
