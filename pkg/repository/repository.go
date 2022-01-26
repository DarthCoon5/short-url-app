package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SaveUrl(longUrl string, shortUrl string) error {
	var exists bool
	existsQuery := fmt.Sprintf("SELECT EXISTS(SELECT $1 FROM %s WHERE long_url = $2)", urlTable)
	existsRow := r.db.QueryRow(existsQuery, longUrl, longUrl)

	if err := existsRow.Scan(&exists); err != nil {
		return err
	} else if !exists {
		insertQuery := fmt.Sprintf("INSERT INTO %s (long_url, short_url) value ($1, $2)", urlTable)
		insetRow := r.db.QueryRow(insertQuery, longUrl, shortUrl)
		if err := insetRow.Scan(&longUrl); err != nil {
			return err
		}
	}

	return nil
}

func (r *Repository) GetLongUrl(shortUrl string) (string, error) {
	var longUrl string
	query := fmt.Sprintf("SELECT long_url FROM %s WHERE short_url = $1", urlTable)
	row := r.db.QueryRow(query, shortUrl)

	if err := row.Scan(&longUrl); err != nil {
		return "", err
	}

	return longUrl, nil
}
