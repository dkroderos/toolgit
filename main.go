package main

import (
	"os"

	"github.com/dkroderos/toolgit/commands"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "toolgit"}

	rootCmd.AddCommand(commands.NewCloneCommand())
	rootCmd.AddCommand(commands.NewSetupCommand())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
