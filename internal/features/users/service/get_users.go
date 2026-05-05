package user_service

import (
	"context"
	"fmt"

	"gitlab.com/voykinEgor/gorestapi/internal/core/domain"
)

func (s *UserService) GetUsers(
	ctx context.Context,
) ([]domain.User, error) {
	users, err := s.userRepository.GetUsers(ctx)

	if err != nil {
		return nil, fmt.Errorf("get users: %w", err)
	}

	return users, nil
}
