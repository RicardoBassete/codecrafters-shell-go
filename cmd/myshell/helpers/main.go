package helpers

import (
	"os"
	"path/filepath"
	"strings"
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
