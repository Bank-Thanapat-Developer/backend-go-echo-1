package usecase

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	_entities "github.com/thanapatjitmung/entities"
	_repository "github.com/thanapatjitmung/repository"
)

type (
	AuthUsecase interface {
		Register(user *_entities.User) error
		Login(user *_entities.User) (*_entities.Token, error)
	}

	authUsecaseImpl struct {
		userRepo _repository.AuthRepo
		jwtKey   []byte
	}
)

func NewAuthUseCaseImpl(userRepo _repository.AuthRepo, jwtKey []byte) AuthUsecase {
	return &authUsecaseImpl{
		userRepo: userRepo,
		jwtKey:   jwtKey,
	}
}

func (a *authUsecaseImpl) Register(user *_entities.User) error {
	data, err := a.userRepo.GetAllData()
	if err != nil {
		return err
	}
	for _, d := range data {
		if user.Username == d.Username {
			return errors.New("username already exists")
		}
	}
	user.ID = generateUniqueID()
	user.Role = "client"
	err = a.userRepo.Save(user)
	return err
}

func (a *authUsecaseImpl) Login(user *_entities.User) (*_entities.Token, error) {
	var token *_entities.Token
	data, err := a.userRepo.GetAllData()
	if err != nil {
		return nil, err
	}
	jwtKeySecret := a.jwtKey

	for _, d := range data {
		if user.Username == d.Username && user.Password == d.Password {
			if d.Role == "admin" {
				jwtKeySecret = []byte("secret-jwt-admin")
			}
			claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"id":       d.ID,
				"username": d.Username,
				"role":     d.Role,
				"exp":      time.Now().Add(time.Hour * 1).Unix(),
			})
			tokenString, err := claims.SignedString(jwtKeySecret)
			if err != nil {
				return nil, err
			}

			token = &_entities.Token{
				Token: tokenString,
			}

			return token, nil
		}
	}

	return nil, errors.New("username or password invalid")
}

func generateUniqueID() int {
	return int(time.Now().UnixNano())
}
