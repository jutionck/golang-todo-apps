package domain

import validation "github.com/go-ozzo/ozzo-validation"

type Todo struct {
	BaseModel
	Name        string `json:"name"`
	IsCompleted bool   `json:"isCompleted"`
	UserID      string
	User        User `gorm:"foreignKey:UserID"`
}

func (u Todo) IsValidField() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required, validation.Length(3, 50)),
	)
}
