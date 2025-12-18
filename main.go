package main

import (
	"bufio"
	"fmt"
	"os"
	exec "os/exec"

	"github.com/google/shlex"
)

func main() {
	for {
		input := getInput()
		parts, _ := shlex.Split(input)
		head := parts[0]
		args := parts[1:]
		cmd := exec.Command(head, args...)

		fmt.Println(head, args)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		cmd.Run()
	}
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	text, _ := reader.ReadString('\n')
	return text
}
