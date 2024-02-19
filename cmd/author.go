package cmd

import "github.com/spf13/cobra"

var authorCmd = &cobra.Command{
	Use:   "author",
	Short: "Author command handler",
	Long:  `handler for the author command. This command is used to create and list authors`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(authorCmd)

	authorCmd.Flags().BoolP("help", "h", false, "learn-cli-go help commands")
}
