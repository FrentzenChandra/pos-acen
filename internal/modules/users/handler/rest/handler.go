package rest

import (
	"encoding/json"
	"net/http"

	"pos-acen/internal/helper"
	"pos-acen/internal/modules/users/entity"
	"pos-acen/internal/modules/users/ports"

	"github.com/thedevsaddam/renderer"
)

type UserHandler struct {
	render  *renderer.Render
	service ports.UserService
}

func NewUserHandler(service ports.UserService, render *renderer.Render) *UserHandler {
	return &UserHandler{
		render:  render,
		service: service,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var bReq entity.User

	if err := json.NewDecoder(r.Body).Decode(&bReq); err != nil {
		helper.HandleResponse(w, h.render, http.StatusConflict, err.Error(), nil)
		return
	}

	bResp, err := h.service.RegisterUser(bReq)

	if err != nil {
		helper.HandleResponse(w, h.render, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.HandleResponse(w, h.render, http.StatusOK, helper.SUCCESS_MESSSAGE, bResp)
}

// Id        uuid.UUID `json:"id" validate:"required"`
// Email     string    `json:"email" validate:"required,email"`
// Username  string    `json:"username" validate:"required"`
// Password  string    `json:"password" validate:"required"`
// CreatedAt string    `json:"created_at"`
// UpdatedAt string    `json:"updated_at"`
// DeletedAt string    `json:"deleted_at"`