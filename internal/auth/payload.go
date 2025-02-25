package auth

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Email string `json:" "`
	Password string `json:"password"`
}
