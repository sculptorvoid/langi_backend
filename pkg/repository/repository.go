package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
