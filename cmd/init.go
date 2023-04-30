package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:    "co-author",
		Long:   "Integrate as git hook and select past contributors to add Co-authored-by entries easily",
		Hidden: true,
	}
)

func init() {
	rootCmd.AddCommand(hookCmd)
	rootCmd.AddCommand(commitCmd)
	rootCmd.AddCommand(versionCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
