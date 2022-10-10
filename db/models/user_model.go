package model

import db "github.com/oddinnovate/a4go/db/sqlc"

type AUTHUser struct {
	ID       int64
	PublicID string
	Username string
	Email    string
}

type UserDTO struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
}

func NewUserDTO(id, username, email, fullName, lastName string) UserDTO {
	return UserDTO{
		ID:       id,
		Username: username,
		Email:    email,
		FullName: fullName,
	}
}

func UserToDTO(user db.User) UserDTO {
	return UserDTO{
		ID:       user.PublicID,
		Username: user.Username,
		Email:    user.Email,
		FullName: user.FullName,
	}
}

type PublicUserDTO struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func NewPublicUserDTO(id, username string) PublicUserDTO {
	return PublicUserDTO{
		ID:       id,
		Username: username,
	}
}
