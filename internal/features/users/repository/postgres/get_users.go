package users_postgres_repository

import (
	"context"
	"fmt"

	"gitlab.com/voykinEgor/gorestapi/internal/core/domain"
)

func (r *UsersRepository) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.GetOperationTimeout())
	defer cancel()

	query := `
	SELECT * 
	FROM todo.users 
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2`

	rows, err := r.pool.Query(ctx, query, limit, offset)

	if err != nil {
		return nil, fmt.Errorf("exec sql script error: %w", err)
	}
	defer rows.Close()

	var usersList []UserModel

	for rows.Next() {
		var userModel UserModel
		if err := rows.Scan(&userModel.ID, &userModel.Version, &userModel.FullName, &userModel.PhoneNumber); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		usersList = append(usersList, userModel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	userDomains := userDomainsFromUserModels(usersList)

	return userDomains, nil
}
