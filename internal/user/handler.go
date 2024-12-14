package user

import "net/http"

type UserHandler struct {
	Service *UserService
}

func NewUserHandler(s *UserService) *UserHandler {
	return &UserHandler{Service: s}
}

// CreateUser handler general para la creación de usuarios
// Los usuarios creados aqui podrán obtener API Key para consumir y crear recursos
// No son usuarios para inscripción o creación de torneos
func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

}
