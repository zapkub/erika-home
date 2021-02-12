package fsutil

import (
	"fmt"
	"os"
	"path"
)

func initHomedir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("cannot access home directory: %+v\n", err)
		os.Exit(2)
	}
	return home
}

var homedir = initHomedir()

func ApplicationDir() string {
	return path.Join(homedir, ".erika")
}
