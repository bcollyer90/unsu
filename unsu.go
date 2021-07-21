package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func run(command string) {
	cmd := exec.Command("bash", "-c", command)
	stdout, _ := cmd.Output()
	fmt.Println(string(stdout))
	return
}

func main() {
	r, _ := regexp.Compile("[sudo]{4}")
	command := strings.Join(os.Args[1:], " ")

	if r.MatchString(command) {
		run(strings.Replace(command, "sudo ", "", -1))
	} else {
		run(command)
	}
}
