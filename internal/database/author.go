package database

import "database/sql"

type Author struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewAuthor(db *sql.DB) *Author {
	return &Author{db: db}
}

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
