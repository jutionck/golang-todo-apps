package dto

type TodoRequestDto struct {
	Name        string `json:"name"`
	IsCompleted bool   `json:"isCompleted"`
}
