package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Dictionary interface {
}

type Word interface {
}

type Repository struct {
	Authorization
	Dictionary
	Word
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
