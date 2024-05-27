package helpers

import (
	"os"
	"strings"
)

func IsOnPath(command string) (string, bool) {
	osPath := os.Getenv("PATH")
	paths := strings.Split(osPath, ":")
	for _, path := range paths {
		if _, err := os.Stat(path + "/" + command); err == nil {
			return path, true
		}
	}
	return "", false
}
