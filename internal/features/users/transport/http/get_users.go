package user_transport_http

import (
	"net/http"

	"gitlab.com/voykinEgor/gorestapi/internal/core/domain"
	core_logger "gitlab.com/voykinEgor/gorestapi/internal/core/logger"
	core_response "gitlab.com/voykinEgor/gorestapi/internal/core/transport/http/response"
)

type UserGetResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func (h *UsersHttpHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_response.NewHTTPResponseHandler(logger, rw)
	logger.Debug("invoke GetUsers handler")

	users, err := h.usersService.GetUsers(ctx)

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get users")
		return
	}

	response := usersDtoFromDomain(users)

	responseHandler.JSONResponse(response, http.StatusOK)
}

func usersDtoFromDomain(users []domain.User) []UserGetResponse {
	result := make([]UserGetResponse, 0, len(users))

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
