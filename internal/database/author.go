package database

import (
	"database/sql"
	"errors"
)

type Author struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewAuthor(db *sql.DB) *Author {
	return &Author{db: db}
}

var (
	ErrAuthorNotFound = errors.New("author not found")
)

func (a *Author) GetAll() ([]Author, error) {
	rows, err := a.db.Query("SELECT id, name, description FROM authors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	authors := []Author{}

	for rows.Next() {
		var author Author

		err := rows.Scan(&author.ID, &author.Name, &author.Description)
		if err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (a *Author) GetById(id string) (Author, error) {
	var author Author
	err := a.db.QueryRow(
		"SELECT id, name, description FROM authors WHERE id = $1",
		id,
	).Scan(&author.ID, &author.Name, &author.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return Author{}, ErrAuthorNotFound
		}

		return Author{}, err
	}

	return author, nil
}
