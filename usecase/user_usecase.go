package usecase

import (
	"fmt"
	"github.com/jutionck/golang-todo-apps/config"
	"github.com/jutionck/golang-todo-apps/domain"
	"github.com/jutionck/golang-todo-apps/repository"
	"github.com/jutionck/golang-todo-apps/utils/commons"
	"github.com/jutionck/golang-todo-apps/utils/exception"
	"github.com/jutionck/golang-todo-apps/utils/model"
)

type UserUseCase interface {
	RegisterNew(payload *domain.User) error
	FindAll(requestQueryParams model.RequestQueryParams) ([]domain.User, model.Paging, error)
	FindById(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindByEmailPassword(email string, password string) (*domain.User, error)
	UpdateData(payload *domain.User) error
	DeleteData(id string) error
	InitData() (string, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func (u *userUseCase) RegisterNew(payload *domain.User) error {
	if err := payload.IsValidField(); err != nil {
		return err
	}
	exists, _ := u.repo.GetEmail(payload.Email)
	if exists.Email == payload.Email {
		return fmt.Errorf("oops, data with email %s have been created", payload.Email)
	}
	password, err := commons.HashPassword(payload.Password)
	exception.CheckError(err)
	payload.Password = password
	payload.Role = config.USER
	return u.repo.Save(payload)
}

func (u *userUseCase) FindAll(requestQueryParams model.RequestQueryParams) ([]domain.User, model.Paging, error) {
	if !requestQueryParams.QueryParams.IsSortValid() {
		return nil, model.Paging{}, fmt.Errorf("oops, invalid sort by: %s", requestQueryParams.QueryParams.Sort)
	}
	return u.repo.List(requestQueryParams)
}

func (u *userUseCase) FindById(id string) (*domain.User, error) {
	user, err := u.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("oops, user with ID %s not found", id)
	}
	return user, nil
}

func (u *userUseCase) FindByEmail(email string) (*domain.User, error) {
	user, err := u.repo.GetEmail(email)
	if err != nil {
		return nil, fmt.Errorf("oops, user with email %s not found", email)
	}
	return user, nil
}

func (u *userUseCase) FindByEmailPassword(email string, password string) (*domain.User, error) {
	user, err := u.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	pwdCheck := commons.CheckPasswordHash(password, user.Password)
	if !pwdCheck {
		return nil, fmt.Errorf("oops, password don't match")
	}
	return user, nil
}

func (u *userUseCase) UpdateData(payload *domain.User) error {
	if err := payload.IsValidField(); err != nil {
		return err
	}
	exists, _ := u.repo.GetEmail(payload.Email)
	if exists.Email == payload.Email && exists.ID != payload.ID {
		return fmt.Errorf("oops, data with email %s have been created", payload.Email)
	}

	if payload.Password != "" {
		password, err := commons.HashPassword(payload.Password)
		exception.CheckError(err)
		payload.Password = password
	}

	return u.repo.Save(payload)
}

func (u *userUseCase) DeleteData(id string) error {
	return u.repo.Delete(id)
}

func (u *userUseCase) InitData() (string, error) {
	return u.repo.Init()
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
