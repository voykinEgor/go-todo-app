package user_service

import (
	"context"

	"gitlab.com/voykinEgor/gorestapi/internal/core/domain"
)

type UserService struct {
	userRepository UserRepository
}

type UserRepository interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)

	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)
}

func NewuserService(
	userRepository UserRepository,
) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}
