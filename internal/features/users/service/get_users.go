package user_service

import (
	"context"
	"fmt"

	"gitlab.com/voykinEgor/gorestapi/internal/core/domain"
	core_errors "gitlab.com/voykinEgor/gorestapi/internal/core/errors"
)

func (s *UserService) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {

	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"limit must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"offset must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	users, err := s.userRepository.GetUsers(ctx, limit, offset)

	if err != nil {
		return nil, fmt.Errorf("get users: %w", err)
	}

	return users, nil
}
