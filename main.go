package main

import (
	_ "embed"
	"fmt"
	"github.com/erkanzileli/co-author/cmd"
	"github.com/erkanzileli/co-author/internal/config"
	"os"
)

func main() {
	if err := config.Init(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
