package domain

import validation "github.com/go-ozzo/ozzo-validation"

type User struct {
	BaseModel
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role"`
	Todos    []Todo `json:"todos,omitempty"`
}

func (u User) IsValidField() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, validation.Length(3, 50)),
		validation.Field(&u.Password, validation.Required, validation.Length(5, 50)),
	)
}
