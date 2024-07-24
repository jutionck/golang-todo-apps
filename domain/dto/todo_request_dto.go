package dto

type TodoRequestDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	IsCompleted bool   `json:"isCompleted"`
}
