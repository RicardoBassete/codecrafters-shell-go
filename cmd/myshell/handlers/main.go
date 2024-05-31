package handlers

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/helpers"
)

type CMD struct {
	// The command that was typed
	Name string
	// Contains the arguments of the command
	Args []string
}

// Executes the built-in "cd" command
func CD(cmd CMD) {
	var target string
	isHome := false

	if len(cmd.Args) == 0 {
		isHome = true // if command has no arguments, set target to home folder
	} else {
		isHome = strings.Split(cmd.Args[0], "")[0] == "~" // check for home character
	}

	if isHome {
		target = os.Getenv("HOME") // get env variable "HOME"
	} else {
		target = cmd.Args[0] // set target to first argument
	}

	err := os.Chdir(target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: No such file or directory\n", target)
	}
}

// Executes the built-in "echo" command
func ECHO(cmd CMD) {
	out := strings.Join(cmd.Args, " ")
	fmt.Fprintf(os.Stdout, "%s\n", out)
}

// Executes the built-in "type" command
func TYPE(cmd CMD) {
	arg := cmd.Args[0]

	if helpers.IsBuiltIn(arg) {
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", arg)
	} else {

		cmdPath, IsOnPath := helpers.IsOnPath(arg)
		if IsOnPath {
			fmt.Fprintf(os.Stdout, "%s is %s\n", arg, cmdPath)
		} else {
			fmt.Fprintf(os.Stderr, "%s not found\n", arg)
		}

	}
}

// Executes the built-in "pwd" command
func PWD(cmd CMD) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Fprintln(os.Stdout, dir)
	}
}

// Executes the built-in "exit" command
func EXIT(cmd CMD) {
	code := 0
	if len(cmd.Args) == 1 {
		var err error
		code, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "exit: %s: numeric argument required\n", cmd.Args[0])
			code = 255
		}
	}
	os.Exit(code)
}
