package user_transport_http

import "gitlab.com/voykinEgor/gorestapi/internal/core/domain"

type UserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func dtoFromDomain(domain domain.User) UserResponse {
	return UserResponse{
		ID:          domain.ID,
		Version:     domain.Version,
		FullName:    domain.FullName,
		PhoneNumber: domain.PhoneNumber,
	}
}

func usersDtoFromDomain(users []domain.User) []UserGetResponse {
	result := make([]UserGetResponse, len(users))

	for _, u := range users {
		dto := UserGetResponse{
			ID:          u.ID,
			Version:     u.Version,
			FullName:    u.FullName,
			PhoneNumber: u.PhoneNumber,
		}
		result = append(result, dto)
	}

	return result
}
