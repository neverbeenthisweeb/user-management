package jsondb

import "usermanagement/model"

type User struct {
	UserID    string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	IsDeleted bool   `json:"is_deleted,omitempty"`
}

func (u User) ID() (string, interface{}) {
	return "id", u.UserID
}

func (u User) FromModel(m model.User) User {
	return User{
		UserID:    m.ID,
		Username:  m.Username,
		Password:  m.Password,
		IsDeleted: m.IsDeleted,
	}
}

func (u User) ToModel() model.User {
	return model.User{
		ID:        u.UserID,
		Username:  u.Username,
		Password:  u.Password,
		IsDeleted: u.IsDeleted,
	}
}
