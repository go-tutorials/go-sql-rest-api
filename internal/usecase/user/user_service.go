package user

import (
	"context"
	sv "github.com/core-go/service"
)

type UserService interface {
	Load(ctx context.Context, id string) (*User, error)
	Insert(ctx context.Context, user *User) (int64, error)
	Update(ctx context.Context, user *User) (int64, error)
	Patch(ctx context.Context, user map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type DefaultUserService struct {
	repository sv.Repository
}

func NewUserService(repository sv.Repository) *DefaultUserService {
	return &DefaultUserService{repository: repository}
}

func (s *DefaultUserService) Load(ctx context.Context, id string) (*User, error) {
	var user User
	ok, err := s.repository.LoadAndDecode(ctx, id, &user)
	if !ok {
		return nil, err
	} else {
		return &user, err
	}
}
func (s *DefaultUserService) Insert(ctx context.Context, user *User) (int64, error) {
	return s.repository.Insert(ctx, user)
}
func (s *DefaultUserService) Update(ctx context.Context, user *User) (int64, error) {
	return s.repository.Update(ctx, user)
}
func (s *DefaultUserService) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, user)
}
func (s *DefaultUserService) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
