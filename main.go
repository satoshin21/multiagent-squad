package main

import (
	"os"

	"github.com/satoshin21/multiagent/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
