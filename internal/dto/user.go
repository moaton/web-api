package dto

import "time"

type CreateUserDTO struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Login     string    `json:"login"`
	UpdatedAt time.Time `json:"updated_at"`
}
