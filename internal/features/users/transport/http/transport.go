package user_transport_http

import (
	"context"
	"net/http"

	"gitlab.com/voykinEgor/gorestapi/internal/core/domain"
	core_server "gitlab.com/voykinEgor/gorestapi/internal/core/transport/http/server"
)

type UsersHttpHandler struct {
	usersService UsersService
}

type UsersService interface {
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

func NewUsersHttpHandler(
	usersService UsersService,
) *UsersHttpHandler {
	return &UsersHttpHandler{usersService: usersService}
}

func (h *UsersHttpHandler) Routes() []core_server.Route {
	return []core_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: h.GetUsers,
		},
	}
}
