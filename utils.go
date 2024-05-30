package main

import (
	"os"
	"strings"
)

func DirExists(path string) bool {
	d, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return d.IsDir()
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func AddTrailingSlash(path string) string {
	if !strings.HasSuffix(path, "/") && path != "" {
		path = path + "/"
	}

	return path
}
