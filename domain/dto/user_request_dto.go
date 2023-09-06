package dto

type UserRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
