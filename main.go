package main

import (
	"os"

	"github.com/cifra-city/distributors-admin/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
