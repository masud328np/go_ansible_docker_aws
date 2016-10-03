package main

import (
	"fmt"
	"os"
)

func main() {
	//	resp, err := http.Get("https://yahoo.com")
	//:	check(err)
	//body, err := ioutil.ReadAll(resp.Body)
	//check(err)
	fmt.Println(1024)
	for {
	}

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
