package main

import (
	"os"

	"github.com/grokify/go-aha/v3/cmd/aha/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
