package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "workmate",
	Short: "Workmate — AI-Powered DevOps & Productivity Agent",
}

func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
