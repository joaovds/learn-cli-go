package cmd

import (
	"github.com/joaovds/learn-cli-go/internal/database"
	"github.com/joaovds/learn-cli-go/internal/service"
	"github.com/spf13/cobra"
)

func newAuthorListCmd(as *service.AuthorService) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List author(s) command",
		Long:  `List author(s) command. This command is used to list all authors or a specific author by id.`,
		RunE:  listAuthor(as),
	}
}

func listAuthor(as *service.AuthorService) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		authors, err := as.GetAuthors()
		if err != nil {
			cmd.PrintErr(err)
			return err
		}

		for _, a := range authors {
			cmd.Println(a.Format())
		}

		return nil
	}
}

func init() {
	db, _ := database.GetConnection()
	authorDB := database.NewAuthor(db)
	authorService := service.NewAuthorService(*authorDB)

	authorCmd.AddCommand(newAuthorListCmd(authorService))
}
