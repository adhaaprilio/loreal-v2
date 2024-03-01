package dto

type RegisterRequest struct {
	Name                 string `json:"name"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
