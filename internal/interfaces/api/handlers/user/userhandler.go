package user

import (
	"encoding/json"
	"jollej/db-scout/internal/application/user"
	userDto "jollej/db-scout/internal/interfaces/api/dto/user"
	"net/http"
)

type UserHandler struct {
	userService *user.UserService
}

func NewUsersHandler(userService *user.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (a *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var usersResponse userDto.ListUsersResponse
	users := a.userService.List(r.Context())
	for _, user := range users {
		usersResponse.Users = append(usersResponse.Users, userDto.ToUserDto(user))
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usersResponse); err != nil {
		http.Error(w, "Error encoding the response", http.StatusInternalServerError)
		return
	}
}
