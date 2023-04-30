package cmd

import (
	"fmt"
	"github.com/erkanzileli/co-author/internal/config"
	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:     "version",
		Short:   "print the version",
		Example: "co-author version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(config.Version)
		},
	}
)
