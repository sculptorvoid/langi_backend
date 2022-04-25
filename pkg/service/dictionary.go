package service

import (
	"github.com/sculptorvoid/langi_backend/pkg/entity"
	"github.com/sculptorvoid/langi_backend/pkg/repository"
)

type DictionaryService struct {
	repo repository.Dictionary
}

func NewDictionaryService(repo repository.Dictionary) *DictionaryService {
	return &DictionaryService{repo: repo}
}

func (s *DictionaryService) CreateDictionary(userId int, dictionary entity.Dictionary) (int, error) {
	return s.repo.CreateDictionary(userId, dictionary)
}

func (s *DictionaryService) GetAllDictionaries(userId int) ([]entity.Dictionary, error) {
	return s.repo.GetAllDictionaries(userId)
}

func (s *DictionaryService) GetById(userId, dictId int) (entity.Dictionary, error) {
	return s.repo.GetById(userId, dictId)
}
