package main

import (
	"os"

	"github.com/satoshin21/multiagent-squad/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
