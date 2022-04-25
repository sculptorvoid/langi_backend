package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sculptorvoid/langi_backend/pkg/entity"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
}

type Dictionary interface {
	CreateDictionary(userId int, dictionary entity.Dictionary) (int, error)
	GetAllDictionaries(userId int) ([]entity.Dictionary, error)
	GetById(userId, dictId int) (entity.Dictionary, error)
	Delete(userId, dictId int) error
	Update(userId, dictId int, input entity.UpdateDictionaryInput) error
}

type Word interface {
	Create(dictId int, word entity.Word) (int, error)
	GetAll(userId, dictId int) ([]entity.Word, error)
	GetById(userId, wordId int) (entity.Word, error)
	Delete(userId, wordId int) error
	Update(userId, wordId int, input entity.UpdateWordInput) error
}

type Repository struct {
	Authorization
	Dictionary
	Word
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Dictionary:    NewDictionaryPsql(db),
		Word:          NewWordPsql(db),
	}
}
