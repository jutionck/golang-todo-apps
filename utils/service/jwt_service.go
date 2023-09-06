package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jutionck/golang-todo-apps/config"
	"github.com/jutionck/golang-todo-apps/domain"
	"github.com/jutionck/golang-todo-apps/utils/model"
	"time"
)

type JwtService interface {
	CreateAccessToken(credential domain.User) (string, error)
	VerifyAccessToken(tokenString string) (jwt.MapClaims, error)
}

type jwtService struct {
	cfg config.TokenConfig
}

func (j *jwtService) CreateAccessToken(credential domain.User) (string, error) {
	now := time.Now().UTC()
	end := now.Add(j.cfg.AccessTokenLifeTime)
	claims := model.MyClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.ApplicationName,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Id:   credential.ID,
		Role: credential.Role,
	}

	jwtNewClaim := jwt.NewWithClaims(j.cfg.JwtSigningMethod, claims)
	token, err := jwtNewClaim.SignedString(j.cfg.JwtSignatureKey)
	if err != nil {
		return "", errors.New("failed to create access token")
	}

	return token, nil
}

func (j *jwtService) VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != j.cfg.JwtSigningMethod {
			return nil, errors.New("signing method invalid")
		}
		return j.cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, errors.New("failed to verify access token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != j.cfg.ApplicationName {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func NewJwtService(cfg config.TokenConfig) JwtService {
	return &jwtService{cfg: cfg}
}
