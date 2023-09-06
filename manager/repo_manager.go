package manager

import "github.com/jutionck/golang-todo-apps/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
	TodoRepo() repository.TodoRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func (r *repoManager) TodoRepo() repository.TodoRepository {
	return repository.NewTodoRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
