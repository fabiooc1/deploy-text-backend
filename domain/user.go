package domain

import "time"

type User struct {
	ID           int64
	Name         string
	Username     string
	PasswordHash string
	RegisteredAt time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
	ProfileImage *File
}
