package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "gitlab.com/voykinEgor/gorestapi/internal/core/logger"
	core_postgres_pool "gitlab.com/voykinEgor/gorestapi/internal/core/repository/postgres/pool"
	core_http_middleware "gitlab.com/voykinEgor/gorestapi/internal/core/transport/http/middleware"
	core_server "gitlab.com/voykinEgor/gorestapi/internal/core/transport/http/server"
	users_postgres_repository "gitlab.com/voykinEgor/gorestapi/internal/features/users/repository/postgres"
	user_service "gitlab.com/voykinEgor/gorestapi/internal/features/users/service"
	user_transport_http "gitlab.com/voykinEgor/gorestapi/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	mainCtx, mainCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer mainCancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())

	if err != nil {
		fmt.Println("Failed to create application logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("initialize postgres pool")
	pool, err := core_postgres_pool.NewConnPool(mainCtx, core_postgres_pool.NewConfigMust())

	if err != nil {
		logger.Fatal("failed to init postgres pool %w", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initialize feture", zap.String("feature", "user"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := user_service.NewuserService(usersRepository)
	usersTransportHTTP := user_transport_http.NewUsersHttpHandler(usersService)

	logger.Debug("initialize HTTP Server")

	httpServer := core_server.NewHttpServer(
		core_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)
	apiVersionRouter := core_server.NewApiVersionRouter(core_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)

	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(mainCtx); err != nil {
		logger.Error("HTTPServer run error:", zap.Error(err))
	}
}
