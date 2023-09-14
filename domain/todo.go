package domain

import validation "github.com/go-ozzo/ozzo-validation"

type Todo struct {
	BaseModel
	Name        string `json:"name"`
	IsCompleted bool   `json:"isCompleted"`
	UserID      string `json:"userId"`
	User        User   `gorm:"foreignKey:UserID" json:"user"`
}

func (u Todo) IsValidField() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required, validation.Length(3, 50)),
	)
}
