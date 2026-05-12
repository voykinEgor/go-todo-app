package user_transport_http

import (
	"fmt"
	"net/http"

	core_logger "gitlab.com/voykinEgor/gorestapi/internal/core/logger"
	core_response "gitlab.com/voykinEgor/gorestapi/internal/core/transport/http/response"
	core_utils "gitlab.com/voykinEgor/gorestapi/internal/core/transport/http/utils"
)

type UserGetResponse UserResponse

func (h *UsersHttpHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_response.NewHTTPResponseHandler(logger, rw)

	limit, offset, err := getLimitOffsetQueryParam(r)

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get 'limit'/'offset' query param")
		return
	}

	users, err := h.usersService.GetUsers(ctx, limit, offset)

	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get users")
		return
	}

	response := usersDtoFromDomain(users)

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParam(r *http.Request) (*int, *int, error) {
	limit, err := core_utils.GetQueryParam(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}

	offset, err := core_utils.GetQueryParam(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}

	return limit, offset, nil
}
