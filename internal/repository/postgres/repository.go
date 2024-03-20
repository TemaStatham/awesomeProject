package postgres

import (
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type Repository struct {
	log *slog.Logger
	db  *sqlx.DB
}

func NewPostgres(db *sqlx.DB, l *slog.Logger) *Repository {
	return &Repository{
		db:  db,
		log: l,
	}
}
