package handlers

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/helpers"
)

type CMD struct {
	Name string
	Args []string
}

func CD(cmd CMD) {
	target := cmd.Args[0]
	err := os.Chdir(target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: No such file or directory\n", target)
	}
}

func ECHO(cmd CMD) {
	out := strings.Join(cmd.Args, " ")
	fmt.Fprintf(os.Stdout, "%s\n", out)
}

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

func PWD(cmd CMD) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Fprintln(os.Stdout, dir)
	}
}

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
