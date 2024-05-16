package service

import (
	"enigma.com/two-gin/model"
	"enigma.com/two-gin/model/dto"
	"enigma.com/two-gin/repository"
)

type tasksService struct {
	repo repository.TasksRepo
}

// CreateTasks implements TasksService.
func (a *tasksService) CreateTasks(tasks model.Task) (model.Task, error) {
	return a.repo.CreateTasks(tasks)
}

// FindAll implements AuthorUseCase.
func (a *tasksService) FindAllTasks(page int, size int) ([]model.Task, dto.Paging, error) {
	return a.repo.FindAllTasks(page, size)
}

// FindById implements AuthorUseCase.
func (a *tasksService) FindByIdTasks(id string) (model.Task, error) {
	return a.repo.FindByIdTasks(id)
}

type TasksService interface {
	FindAllTasks(page int, size int) ([]model.Task, dto.Paging, error)
	FindByIdTasks(id string) (model.Task, error)
	CreateTasks(model.Task) (model.Task, error)
}

func NewTasksService(repo repository.TasksRepo) TasksService {
	return &tasksService{repo: repo}
}
