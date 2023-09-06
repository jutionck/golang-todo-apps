package dto

type LoginRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
