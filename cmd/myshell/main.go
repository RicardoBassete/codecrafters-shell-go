package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
	}

	run(strings.TrimRight(input, "\n"))
}

func run(command string) {
	fmt.Fprintf(os.Stderr, "%s: command not found\n", command)
}
