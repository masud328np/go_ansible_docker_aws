package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

var (
	imgId       string
	containerId string
)

func TestMain(m *testing.M) {
	StartContainer()
	m.Run()

	StopContainer(strings.TrimSpace(containerId), strings.TrimSpace(imgId))
	//os.Exit(v)
}

/* private function*/
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
	result, err := ExecCommand("docker", []string{"build", "-t", "tmp:latest", "-f", "../docker/Dockerfile", "-q", "./../"})
	imageId = string(result)
	return
}

func RunContainer(imgId string) (contId string, err error) {

	fmt.Println("2. Running Container:", containerId)
	result, err := ExecCommand("docker", []string{"run", "-d", "-p", "8080:8080", imgId})
	contId = string(result)
	return
}

func ExecCommand(command string, cmdArgs []string) (output []byte, err error) {

	output, err = exec.Command(command, cmdArgs...).Output()

	/*
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
			return err
		}
	*/
	/*
		scanner := bufio.NewScanner(stdout)
		go func() {
			for scanner.Scan() {
				fmt.Printf("%s", scanner.Text())
			}
		}()
	*/
	/*
		if err := cmd.CombinedOutput; err != nil {
			log.Fatal(err)
		}
	*/

	/*	err := cmd.Start()

		if err != nil {
			log.Fatal(err)
			return err
		}
		if err := cmd.Wait(); err != nil {
			log.Fatal(err)
			return err
		}
	*/
	//fmt.Println("%s\n", string(output))
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
