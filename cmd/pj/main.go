package main

import (
	"os"

	"github.com/lcrownover/pj/internal/cli"
)

func main() {
	cli.Cli(os.Args[1:])
}
