package service

import (
	"github.com/sculptorvoid/langi_backend/pkg/entity"
	"github.com/sculptorvoid/langi_backend/pkg/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Dictionary interface {
	CreateDictionary(userId int, dictionary entity.Dictionary) (int, error)
	GetAllDictionaries(userId int) ([]entity.Dictionary, error)
	GetById(userId, dictId int) (entity.Dictionary, error)
}

type Word interface {
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
	}
}
