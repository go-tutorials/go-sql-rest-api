package app

import (
	"context"
	"github.com/core-go/health"
	"github.com/core-go/log"
	sv "github.com/core-go/service"
	v "github.com/core-go/service/v10"
	q "github.com/core-go/sql"
	"github.com/core-go/sql/query"
	_ "github.com/go-sql-driver/mysql"
	"reflect"

	"go-service/internal/usecase/user"
)

type ApplicationContext struct {
	HealthHandler *health.Handler
	UserHandler   user.UserHandler
}

func NewApp(ctx context.Context, root Root) (*ApplicationContext, error) {
	db, err := q.OpenByConfig(root.Sql)
	if err != nil {
		return nil, err
	}
	logError := log.ErrorMsg
	status := sv.InitializeStatus(root.Status)

	userType := reflect.TypeOf(user.User{})
	userQueryBuilder := query.NewBuilder(db, "users", userType)
	userSearchBuilder, err := q.NewSearchBuilder(db, userType, userQueryBuilder.BuildQuery)
	if err != nil {
		return nil, err
	}
	userRepository, err := q.NewRepository(db, "users", userType)
	if err != nil {
		return nil, err
	}
	userService := user.NewUserService(userRepository)
	validator := v.NewValidator()
	userHandler := user.NewUserHandler(userSearchBuilder.Search, userService, status, validator.Validate, logError)

	sqlChecker := q.NewHealthChecker(db)
	healthHandler := health.NewHandler(sqlChecker)

	return &ApplicationContext{
		HealthHandler: healthHandler,
		UserHandler:   userHandler,
	}, nil
}
