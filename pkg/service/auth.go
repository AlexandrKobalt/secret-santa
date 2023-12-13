package service

import (
	"secret-santa/pkg"
	"secret-santa/pkg/repository"
)

type AuthService struct {
	repository repository.Authentication
}

func NewAuthService(repository repository.Authentication) *AuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) IsUserExists(user pkg.User) (bool, error) {
	return s.repository.IsUserExists(user)
}

func (s *AuthService) AddUser(user pkg.User) (int, error) {

}
