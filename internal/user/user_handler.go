package user

import (
	"context"
	"net/http"
	"reflect"

	sv "github.com/core-go/service"
)

type UserHandler struct {
	service   UserService
	keys      []string
	indexes   map[string]int
	modelType reflect.Type
	Status    sv.StatusConfig
	Validate  func(ctx context.Context, model interface{}) ([]sv.ErrorMessage, error)
	Error     func(context.Context, string)
}

func NewUserHandler(service UserService, status sv.StatusConfig, validate func(ctx context.Context, model interface{}) ([]sv.ErrorMessage, error), logError func(context.Context, string)) *UserHandler {
	modelType := reflect.TypeOf(User{})
	keys, indexes, _ := sv.BuildHandlerParams(modelType)
	return &UserHandler{service: service, keys: keys, indexes: indexes, modelType: modelType, Status: status, Validate: validate, Error: logError}
}

func (h *UserHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		result, err := h.service.Load(r.Context(), id)
		sv.RespondModel(w, r, result, err, h.Error, nil)
	}
}
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	h.Insert(w, r)
}
func (h *UserHandler) Insert(w http.ResponseWriter, r *http.Request) {
	var user User
	er1 := sv.Decode(w, r, &user)
	if er1 == nil {
		errors, er2 := h.Validate(r.Context(), &user)
		if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, nil) {
			result, er3 := h.service.Insert(r.Context(), &user)
			sv.AfterCreated(w, r, &user, result, er3, h.Status, h.Error, nil)
		}
	}
}
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var user User
	er1 := sv.Decode(w, r, &user)
	if er1 == nil {
		er2 := sv.CheckId(w, r, &user, h.keys, h.indexes)
		if er2 == nil {
			errors, er3 := h.Validate(r.Context(), &user)
			if !sv.HasError(w, r, errors, er3, *h.Status.ValidationError, h.Error, nil) {
				result, er4 := h.service.Update(r.Context(), &user)
				sv.HandleResult(w, r, &user, result, er4, h.Status, h.Error, nil)
			}
		}
	}
}
func (h *UserHandler) Patch(w http.ResponseWriter, r *http.Request) {
	r = r.WithContext(context.WithValue(r.Context(), sv.Method, sv.Patch))
	var user User
	body, er0 := sv.BuildMapAndStruct(r, &user, w)
	if er0 == nil {
		er1 := sv.CheckId(w, r, &user, h.keys, h.indexes)
		if er1 == nil {
			errors, er2 := h.Validate(r.Context(), &user)
			if !sv.HasError(w, r, errors, er2, *h.Status.ValidationError, h.Error, nil) {
				json, er3 := sv.BodyToJson(w, r, user, body, h.keys, h.indexes, nil, h.Error, nil)
				if er3 != nil {
					http.Error(w, er3.Error(), http.StatusInternalServerError)
					return
				}
				result, er4 := h.service.Patch(r.Context(), json)
				sv.HandleResult(w, r, json, result, er4, h.Status, h.Error, nil)
			}
		}
	}
}
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := sv.GetRequiredParam(w, r)
	if len(id) > 0 {
		result, err := h.service.Delete(r.Context(), id)
		sv.HandleDelete(w, r, result, err, h.Error, nil)
	}
}
