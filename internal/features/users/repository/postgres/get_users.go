package users_postgres_repository

import (
	"context"
	"fmt"

	"gitlab.com/voykinEgor/gorestapi/internal/core/domain"
)

func (r *UsersRepository) GetUsers(
	ctx context.Context,
) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.GetOperationTimeout())
	defer cancel()

	query := `SELECT * FROM todo.users`

	rows, err := r.pool.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("exec sql script error: %w", err)
	}
	defer rows.Close()

	var usersList []domain.User

	for rows.Next() {
		var id, version int
		var full_name string
		var phone_number *string

		if err := rows.Scan(&id, &version, &full_name, &phone_number); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		user := domain.User{
			ID:          id,
			FullName:    full_name,
			Version:     version,
			PhoneNumber: phone_number,
		}

		usersList = append(usersList, user)
	}

	return usersList, nil
}
