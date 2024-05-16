package service

import (
	"enigma.com/two-gin/model"
	"enigma.com/two-gin/model/dto"
	"enigma.com/two-gin/repository"
)

type authorUseCase struct {
	repo repository.AuthorRepo
}

// FindAll implements AuthorUseCase.
func (a *authorUseCase) FindAll(page int, size int) ([]model.Author, dto.Paging, error) {
	return a.repo.FindAll(page, size)
}

// FindById implements AuthorUseCase.
func (a *authorUseCase) FindById(id string) (model.Author, error) {
	return a.repo.FindById(id)
}

type AuthorUseCase interface {
	FindAll(page int, size int) ([]model.Author, dto.Paging, error)
	FindById(id string) (model.Author, error)
}

func NewAuthorUseCase(repo repository.AuthorRepo) AuthorUseCase {
	return &authorUseCase{repo: repo}
}
