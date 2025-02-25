package auth

import (
	"fmt"
	"go/api-demo/configs"
	"go/api-demo/pkg/resp"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}
type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := LoginResponse{
			Token: "678",
		}
		resp.WriteJson(w, data, 200)
	}
}

func (handelr *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")
	}
}
