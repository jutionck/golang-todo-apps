package dto

type UserRequestDto struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
