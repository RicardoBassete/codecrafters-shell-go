package helpers

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	BUILTIN_COMMANDS = []string{"echo", "pwd", "type", "exit"}
)

func IsOnPath(command string) (string, bool) {
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

func IsBuiltIn(command string) bool {
	for _, builtin := range BUILTIN_COMMANDS {
		if command == builtin {
			return true
		}
	}
	return false
}
