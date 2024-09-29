package httphandlers

import (
	"ddd-template/usecases"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type UserHandler struct {
	UseCase *usecases.UserUseCase
}

func NewUserHandler(u *usecases.UserUseCase) *UserHandler {
	return &UserHandler{
		UseCase: u,
	}
}

func (h *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	if userId == "" {
		w.Write([]byte(fmt.Sprintf("err %v", "missing userId")))
		return
	}
	usr, err := h.UseCase.GetUserById(r.Context(), userId)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err %v", err)))
		return
	}
	w.Write([]byte(fmt.Sprintf("%v", usr)))
}
