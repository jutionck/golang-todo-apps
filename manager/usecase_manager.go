package manager

import "github.com/jutionck/golang-todo-apps/usecase"

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
	TodoUseCase() usecase.TodoUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repo.UserRepo())
}

func (u *useCaseManager) TodoUseCase() usecase.TodoUseCase {
	return usecase.NewTodoUseCase(u.repo.TodoRepo())
}

func NewUseCaseManager(rm RepoManager) UseCaseManager {
	return &useCaseManager{repo: rm}
}
