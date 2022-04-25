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

func (s *DictionaryService) Delete(userId, dictId int) error {
	return s.repo.Delete(userId, dictId)
}

func (s *DictionaryService) Update(userId, dictId int, input entity.UpdateDictionaryInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, dictId, input)
}
