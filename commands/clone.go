package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCloneCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "cl",
		Short: "Clone a repository",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Cloning repository...")
		},
	}
}
