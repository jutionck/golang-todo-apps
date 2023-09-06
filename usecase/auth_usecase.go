package usecase

import (
	"github.com/jutionck/golang-todo-apps/config"
	"github.com/jutionck/golang-todo-apps/domain"
	"github.com/jutionck/golang-todo-apps/utils/service"
)

type AuthenticationUseCase interface {
	Login(username string, password string) (string, error)
	Register(payload *domain.User) error
}

type authenticationUseCase struct {
	userUseCase UserUseCase
	jwtService  service.JwtService
}

func (a *authenticationUseCase) Register(payload *domain.User) error {
	err := payload.IsValidField()
	if err != nil {
		return err
	}
	payload.Role = config.USER
	return a.userUseCase.RegisterNew(payload)
}

func (a *authenticationUseCase) Login(username string, password string) (string, error) {
	user, err := a.userUseCase.FindByEmailPassword(username, password)
	var token string
	if err != nil {
		return "", err
	}
	if user != nil {
		token, err = a.jwtService.CreateAccessToken(*user)
		if err != nil {
			return "", err
		}
	}
	return token, nil
}

func NewAuthenticationUseCase(uc UserUseCase, jwtService service.JwtService) AuthenticationUseCase {
	return &authenticationUseCase{userUseCase: uc, jwtService: jwtService}
}
