package main

import (
	_ "embed"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	//go:embed prepare-commit-msg.template
	prepareCommitMsgHookTemplate string
)

var (
	rootCmd = &cobra.Command{
		Use:     "co-author",
		Long:    "Integrate as git hook and select past contributors to add Co-authored-by entries easily",
		Version: version,
		Hidden:  true,
	}
	version string
)

func init() {
	rootCmd.AddCommand(hookCmd)
	rootCmd.AddCommand(commitCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
