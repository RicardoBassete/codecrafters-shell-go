package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/helpers"
)

type cmd struct {
	name string
	args []string
}

func parseCmd(str string) cmd {
	parts := strings.Fields(str)
	name := parts[0]
	args := parts[1:]
	return cmd{name, args}
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

func run(cmd cmd) {
	switch cmd.name {
	case "echo":
		handleEcho(cmd)
	case "pwd":
		handlePwd(cmd)
	case "type":
		handleType(cmd)
	case "exit":
		handleExit(cmd)
	default:
		command := exec.Command(cmd.name, cmd.args...)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err := command.Run()
		if err != nil {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd.name)
		}
	}
}

func handlePwd(cmd cmd) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
	} else {
		fmt.Fprintln(os.Stdout, dir)
	}
}

func handleType(cmd cmd) {
	arg := cmd.args[0]
	switch arg {
	case "echo":
		fmt.Fprintln(os.Stdout, "echo is a shell builtin")
	case "pwd":
		fmt.Fprintln(os.Stdout, "pwd is a shell builtin")
	case "type":
		fmt.Fprintln(os.Stdout, "type is a shell builtin")
	case "exit":
		fmt.Fprintln(os.Stdout, "exit is a shell builtin")
	default:
		cmdPath, IsOnPath := helpers.IsOnPath(arg)
		if IsOnPath {
			fmt.Fprintf(os.Stdout, "%s is %s\n", arg, cmdPath)
		} else {
			fmt.Fprintf(os.Stdout, "%s not found\n", arg)
		}
	}
}

func handleExit(cmd cmd) {
	code := 0
	if len(cmd.args) == 1 {
		var err error
		code, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			fmt.Fprintf(os.Stdout, "exit: %s: numeric argument required\n", cmd.args[0])
			code = 255
		}
	}
	os.Exit(code)
}

func handleEcho(cmd cmd) {
	out := strings.Join(cmd.args, " ")
	fmt.Fprintf(os.Stdout, "%s\n", out)
}
