package users

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Register(
	ctx context.Context,
	email string,
	password string,
) (*User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, fmt.Errorf("hash password %w:", err)
	}

	user, err := s.repository.Create(
		ctx,
		email,
		string(passwordHash),
	)
	if err != nil {
		return nil, fmt.Errorf("register user %w:", err)
	}

	return user, nil
}
