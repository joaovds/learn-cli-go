package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List author(s) command",
	Long:  `List author(s) command. This command is used to list all authors or a specific author by id.`,
	RunE:  listAuthor(),
}

func listAuthor() RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Println("list called")

		return nil
	}
}

func init() {
	authorCmd.AddCommand(listCmd)
}
