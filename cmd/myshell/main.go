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

func handleExit(cmd cmd) {
	if cmd.name != "exit" {
		return
	}
	code := 0
	if len(cmd.args) == 1 {
		var err error
		code, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			fmt.Printf("exit: %s: numeric argument required\n", cmd.args[0])
			code = 255
		}
	}
	os.Exit(code)
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
	handleExit(cmd)

	fmt.Fprintf(os.Stderr, "%s: command not found\n", cmd.name)
}
