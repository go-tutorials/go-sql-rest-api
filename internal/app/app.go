package app

import (
	"context"
	"github.com/core-go/health"
	"github.com/core-go/log"
	sv "github.com/core-go/service"
	h "github.com/core-go/service/handler"
	v "github.com/core-go/service/v10"
	"github.com/core-go/sql"
	_ "github.com/go-sql-driver/mysql"
	"reflect"

	"go-service/internal/user"
)

type ApplicationContext struct {
	HealthHandler *health.Handler
	UserHandler   h.Handler
}

func NewApp(context context.Context, root Root) (*ApplicationContext, error) {
	db, err := sql.OpenByConfig(root.Sql)
	if err != nil {
		return nil, err
	}
	logError := log.ErrorMsg
	status := sv.InitializeStatus(root.Status)
	userType := reflect.TypeOf(user.User{})
	userRepository, err := sql.NewRepository(db, "users", userType)
	if err != nil {
		return nil, err
	}
	userService := user.NewUserService(userRepository)
	validator := v.NewValidator()
	userHandler := user.NewUserHandler(userService, status, validator.Validate, logError)

	sqlChecker := sql.NewHealthChecker(db)
	healthHandler := health.NewHandler(sqlChecker)

	return &ApplicationContext{
		HealthHandler: healthHandler,
		UserHandler:   userHandler,
	}, nil
}
