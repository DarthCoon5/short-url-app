package repository

import (
	"fmt"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
)

const (
	urlTable = "url"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	SslMode  string
}

func NewPostgresDb(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s username=%spassword=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DbName, cfg.SslMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
