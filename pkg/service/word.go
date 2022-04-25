package service

import (
	"github.com/sculptorvoid/langi_backend/pkg/entity"
	"github.com/sculptorvoid/langi_backend/pkg/repository"
)

type WordService struct {
	repo     repository.Word
	dictRepo repository.Dictionary
}

func NewWordService(repo repository.Word, dictRepo repository.Dictionary) *WordService {
	return &WordService{repo: repo, dictRepo: dictRepo}
}

func (s *WordService) Create(userId, dictId int, word entity.Word) (int, error) {
	_, err := s.dictRepo.GetById(userId, dictId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(dictId, word)
}

func (s *WordService) GetAll(userId, dictId int) ([]entity.Word, error) {
	return s.repo.GetAll(userId, dictId)
}

func (s *WordService) GetById(userId, wordId int) (entity.Word, error) {
	return s.repo.GetById(userId, wordId)
}

func (s *WordService) Delete(userId, wordId int) error {
	return s.repo.Delete(userId, wordId)
}

func (s *WordService) Update(userId, wordId int, input entity.UpdateWordInput) error {
	return s.repo.Update(userId, wordId, input)
}
