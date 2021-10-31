package user

import (
	"context"
	"database/sql"
	sv "github.com/core-go/service"
	q "github.com/core-go/sql"
)

type UserService interface {
	Load(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *User) (int64, error)
	Update(ctx context.Context, user *User) (int64, error)
	Patch(ctx context.Context, user map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewUserService(db *sql.DB, repository sv.Repository) UserService {
	return &userService{db: db, repository: repository}
}

type userService struct {
	db         *sql.DB
	repository sv.Repository
}

func (s *userService) Load(ctx context.Context, id string) (*User, error) {
	var user User
	ok, err := s.repository.LoadAndDecode(ctx, id, &user)
	if !ok {
		return nil, err
	} else {
		return &user, err
	}
}
func (s *userService) Create(ctx context.Context, user *User) (int64, error) {
	ctx, tx, err := q.Begin(ctx, s.db)
	if err != nil {
		return  -1, err
	}
	res, err := s.repository.Insert(ctx, user)
	return q.End(tx, res, err)
}
func (s *userService) Update(ctx context.Context, user *User) (int64, error) {
	ctx, tx, err := q.Begin(ctx, s.db)
	if err != nil {
		return  -1, err
	}
	res, err := s.repository.Update(ctx, user)
	err = q.Commit(tx, err)
	return res, err
}
func (s *userService) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, user)
}
func (s *userService) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
