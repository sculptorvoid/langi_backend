package service

import "github.com/sculptorvoid/langi_backend/pkg/repository"

type Authorization interface {
}

type Dictionary interface {
}

type Word interface {
}

type Service struct {
	Authorization
	Dictionary
	Word
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
