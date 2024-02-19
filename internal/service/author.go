package service

import "github.com/joaovds/learn-cli-go/internal/database"

type AuthorServiceResponse struct {
	ID          string
	Name        string
	Description string
}

func NewAuthorServiceResponse(id, name, description string) *AuthorServiceResponse {
	return &AuthorServiceResponse{
		ID:          id,
		Name:        name,
		Description: description,
	}
}

func (a *AuthorServiceResponse) Format() string {
	return `ID: ` + a.ID + ` | Name: ` + a.Name + ` | Description: ` + a.Description
}

type AuthorService struct {
	AuthorDB database.Author
}

func NewAuthorService(authorDB database.Author) *AuthorService {
	return &AuthorService{
		AuthorDB: authorDB,
	}
}

func (a *AuthorService) GetAuthors() ([]AuthorServiceResponse, error) {
	authors, err := a.AuthorDB.GetAll()
	if err != nil {
		return nil, err
	}

	response := make([]AuthorServiceResponse, len(authors))

	for i, author := range authors {
		response[i] = *NewAuthorServiceResponse(author.ID, author.Name, author.Description)
	}

	return response, nil
}
