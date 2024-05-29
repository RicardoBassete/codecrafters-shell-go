package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/handlers"
)

func parseCmd(str string) handlers.CMD {
	parts := strings.Fields(str)
	name := parts[0]
	args := parts[1:]
	return handlers.CMD{Name: name, Args: args}
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}
		cmd := parseCmd(input[:len(input)-1])
		run(cmd)
	}
}

func run(cmd handlers.CMD) {
	switch cmd.Name {
	case "cd":
		handlers.CD(cmd)
	case "echo":
		handlers.ECHO(cmd)
	case "pwd":
		handlers.PWD(cmd)
	case "type":
		handlers.TYPE(cmd)
	case "exit":
		handlers.EXIT(cmd)
	default:
		command := exec.Command(cmd.Name, cmd.Args...)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err := command.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: command not found\n", cmd.Name)
		}
	}
}
