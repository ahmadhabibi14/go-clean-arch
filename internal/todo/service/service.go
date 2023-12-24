package service

import "go-clean-arch/internal/todo/repository"

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
		repo: r,
	}
}
