package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewSetupCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "setup",
		Short: "Setup toolgit",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Setting up...")
		},
	}
}
