package auth_test

import (
	"bytes"
	"encoding/json"
	"go/api-demo/configs"
	"go/api-demo/internal/auth"
	"go/api-demo/internal/user"
	"go/api-demo/pkg/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func bootstrap() (*auth.AuthHandler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,  
	}))
	if err != nil {
		return nil, nil, err
	}
	userRepo := user.NewUserRepository(&db.Db{
		DB: gormDb,
	})
	handler := auth.AuthHandler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		AuthService: auth.NewAuthService(userRepo),
	}
	return &handler, mock, nil
}

func TestLoginSuccess(t *testing.T) {
	handler, _, err := bootstrap()
	if err != nil {
		t.Fatal(err)
		return
	}
	data, _ := json.Marshal(&auth.LoginRequest{
		Email: "almaz@yandex.ru",
		Password: "123456789",
	})
	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", reader)
	handler.Login()(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("got %d, expected %d", w.Code, 200)
	}
}
