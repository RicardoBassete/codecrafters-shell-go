package helpers

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	// List of all available built-in commands
	BUILTIN_COMMANDS = []string{"cd", "echo", "pwd", "type", "exit"}
)

// Check if the command is on $PATH
func IsOnPath(command string) (path string, is_on_path bool) {
	osPath := os.Getenv("PATH")
	paths := strings.Split(osPath, ":")
	for _, path := range paths {
		fullpath := filepath.Join(path, command)
		if _, err := os.Stat(fullpath); err == nil {
			return fullpath, true
		}
	}
	return "", false
}

// Check if the command is a built-in command
func IsBuiltIn(command string) bool {
	for _, builtin := range BUILTIN_COMMANDS {
		if command == builtin {
			return true
		}
	}
	return false
}
