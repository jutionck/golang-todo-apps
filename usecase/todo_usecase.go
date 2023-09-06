package usecase

import (
	"fmt"
	"github.com/jutionck/golang-todo-apps/domain"
	"github.com/jutionck/golang-todo-apps/repository"
	"github.com/jutionck/golang-todo-apps/utils/model"
)

type TodoUseCase interface {
	RegisterNew(payload *domain.Todo) error
	FindAll(requestQueryParams model.RequestQueryParams, excludeUserID string) ([]domain.Todo, model.Paging, error)
	FindById(id string) (*domain.Todo, error)
	UpdateData(payload *domain.Todo) error
	DeleteData(id string) error
}

type todoUseCase struct {
	repo repository.TodoRepository
}

func (t *todoUseCase) RegisterNew(payload *domain.Todo) error {
	if err := payload.IsValidField(); err != nil {
		return err
	}
	return t.repo.Save(payload)
}

func (t *todoUseCase) FindAll(requestQueryParams model.RequestQueryParams, excludeUserID string) ([]domain.Todo, model.Paging, error) {
	if !requestQueryParams.QueryParams.IsSortValid() {
		return nil, model.Paging{}, fmt.Errorf("oops, invalid sort by: %s", requestQueryParams.QueryParams.Sort)
	}
	return t.repo.List(requestQueryParams, excludeUserID)
}

func (t *todoUseCase) FindById(id string) (*domain.Todo, error) {
	todo, err := t.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("oops, todo with ID %s not found", id)
	}
	return todo, nil
}

func (t *todoUseCase) UpdateData(payload *domain.Todo) error {
	if err := payload.IsValidField(); err != nil {
		return err
	}
	return t.repo.Save(payload)
}

func (t *todoUseCase) DeleteData(id string) error {
	return t.repo.Delete(id)
}

func NewTodoUseCase(repo repository.TodoRepository) TodoUseCase {
	return &todoUseCase{repo: repo}
}
