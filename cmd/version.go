package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version   = "dev"
	Commit    = "none"
	GoVersion = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show build version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version:    %s\n", Version)
		fmt.Printf("commit:     %s\n", Commit)
		fmt.Printf("go version: %s\n", GoVersion)
	},
}

func init() {
	Cmd.AddCommand(versionCmd)
}
