package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	case "exit":
		handleExit(cmd)
	case "echo":
		handleEcho(cmd)
	default:
		fmt.Fprintf(os.Stderr, "%s: command not found\n", cmd.name)
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
