package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {

	r, _ := regexp.Compile("[sudo]{4}")

	command := strings.Join(os.Args[1:], " ")

	if r.MatchString(command) {
		cmd := exec.Command("bash", "-c", strings.Replace(command, "sudo ", "", -1))
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(string(stdout))
	} else {
		cmd := exec.Command("bash", "-c", command)
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(string(stdout))
	}
}
