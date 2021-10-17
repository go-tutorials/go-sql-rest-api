package app

import (
	"context"
	"github.com/core-go/health"
	s "github.com/core-go/health/sql"
	"github.com/core-go/sql"
	_ "github.com/go-sql-driver/mysql"

	"go-service/internal/user"
)

type ApplicationContext struct {
	HealthHandler *health.Handler
	UserHandler   *user.UserHandler
}

func NewApp(context context.Context, root Root) (*ApplicationContext, error) {
	db, err := sql.OpenByConfig(root.Sql)
	if err != nil {
		return nil, err
	}

	userService := user.NewUserService(db)
	userHandler := user.NewUserHandler(userService)

	sqlChecker := s.NewHealthChecker(db)
	healthHandler := health.NewHandler(sqlChecker)

	return &ApplicationContext{
		HealthHandler: healthHandler,
		UserHandler:   userHandler,
	}, nil
}
