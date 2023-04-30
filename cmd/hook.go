package cmd

import (
	"fmt"
	"github.com/erkanzileli/co-author/internal/git"
	"github.com/spf13/cobra"
)

var (
	hookCmd = &cobra.Command{
		Use:     "hook",
		Short:   "print a hook configuration",
		Long:    "this prints a git hook",
		Example: "co-author hook > .git/hooks/prepare-commit-msg",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(git.PrepareCommitMsgHookTemplate)
		},
	}
)
