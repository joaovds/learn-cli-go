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
		authorID, _ := cmd.Flags().GetString("id")

		if authorID != "" {
			return listAuthorById(cmd, as, authorID)
		}

		return listAllAuthors(cmd, as)
	}
}

func listAllAuthors(cmd *cobra.Command, as *service.AuthorService) error {
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

func listAuthorById(cmd *cobra.Command, as *service.AuthorService, id string) error {
	author, err := as.GetAuthorById(id)
	if err != nil {
		cmd.PrintErr(err)
		return err
	}

	cmd.Println(author.Format())

	return nil
}

func init() {
	db, _ := database.GetConnection()
	authorDB := database.NewAuthor(db)
	authorService := service.NewAuthorService(*authorDB)

	listCmd := newAuthorListCmd(authorService)

	authorCmd.AddCommand(listCmd)

	listCmd.Flags().String("id", "", "Author ID")
}
