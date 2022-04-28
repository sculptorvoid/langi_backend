package service

import (
	"github.com/sculptorvoid/langi_backend/pkg/entity"
	"github.com/sculptorvoid/langi_backend/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Dictionary interface {
	CreateDictionary(userId int, dictionary entity.Dictionary) (int, error)
	GetAllDictionaries(userId int) ([]entity.Dictionary, error)
	GetById(userId, dictId int) (entity.Dictionary, error)
	Delete(userId, dictId int) error
	Update(userId, dictId int, input entity.UpdateDictionaryInput) error
}

type Word interface {
	Create(userId, dictId int, word entity.Word) (int, error)
	GetAll(userId, dictId int) ([]entity.Word, error)
	GetById(userId, wordId int) (entity.Word, error)
	Delete(userId, wordId int) error
	Update(userId, wordId int, input entity.UpdateWordInput) error
}

type Service struct {
	Authorization
	Dictionary
	Word
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Dictionary:    NewDictionaryService(repos.Dictionary),
		Word:          NewWordService(repos.Word, repos.Dictionary),
	}
}
