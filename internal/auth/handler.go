package auth

import (
	"fmt"
	"go/api-demo/configs"
	"go/api-demo/pkg/req"
	"go/api-demo/pkg/resp"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService 
}

type AuthHandler struct {
	*configs.Config
	*AuthService
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
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
		data := LoginResponse{
			Token: "678",
		}
		resp.WriteJson(w, data, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		handler.AuthService.Register(body.Email, body.Password, body.Name)
		resp.WriteJson(w, body, 200)
	}
}
