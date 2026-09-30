package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mirily/shorty/internal/users"
)

type AuthHandler struct {
	userService *users.Service
}

func NewAuthHandler(userService *users.Service) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerResponse struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

func (h *AuthHandler) Register(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req registerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)
		return
	}

	user, err := h.userService.Register(
		r.Context(),
		req.Email,
		req.Password,
	)

	if err != nil {
		http.Error(
			w,
			"failed to register user",
			http.StatusInternalServerError,
		)

		return
	}

	response := registerResponse{
		ID:    user.ID,
		Email: user.Email,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		return
	}
}
