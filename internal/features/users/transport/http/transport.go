package user_transport_http

type UsersHttpHandler struct {
	usersService UsersService
}

type UsersService interface {
}

func NewUsersHttpHandler(
	usersService UsersService,
) *UsersHttpHandler {
	return &UsersHttpHandler{usersService: usersService}
}
