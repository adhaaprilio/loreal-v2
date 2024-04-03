package entity

type User struct {
	Id       string `db:"id"`
	Name     string `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type RegisterResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	// AccessToken string `json:"accessToken"`
}

type ResponseParams struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
