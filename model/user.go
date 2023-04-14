package model

type User struct {
	ID        string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	IsDeleted bool   `json:"is_deleted"`
}

type UserListFilter struct {
	ShowDeleted bool
}
