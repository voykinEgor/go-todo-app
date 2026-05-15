package user_transport_http

import (
	"net/http"

	"gitlab.com/voykinEgor/gorestapi/internal/core/domain"
	core_logger "gitlab.com/voykinEgor/gorestapi/internal/core/logger"
	core_http_request "gitlab.com/voykinEgor/gorestapi/internal/core/transport/http/request"
	core_response "gitlab.com/voykinEgor/gorestapi/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+"`
}

type CreateUserResponse UserResponse

func (h *UsersHttpHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_response.NewHTTPResponseHandler(logger, rw)
	logger.Debug("invoke CreateUser handler")

	var request CreateUserRequest
	if err := core_http_request.Decode(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "decode and validate error")
		return
	}
	userDomain := domainFromDTO(request)
	user, err := h.usersService.CreateUser(ctx, userDomain)

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	response := CreateUserResponse(dtoFromDomain(user))

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.FullName, dto.PhoneNumber)
}
