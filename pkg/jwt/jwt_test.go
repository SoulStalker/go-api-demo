package jwt_test

import (
	"go/api-demo/pkg/jwt"
	"testing"
)

func TestJWTCreate(t *testing.T) {
	const email = "almaz@ya.ru"
	jwtService := jwt.NewJWT("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid{
		t.Fatal("Token is invalid")
	}
	if data.Email != email {
		t.Fatalf("Email %s not equal %s", data.Email, email)
	}
}
