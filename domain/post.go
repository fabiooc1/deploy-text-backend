package domain

import "time"

type Post struct {
	ID            int64
	Title         string
	Slug          string
	Body          string
	ViewsCount    int64
	FavoriteCount int64
	PostedAt      time.Time
	UpdatedAt     *time.Time
	DeletedAt     *time.Time
	PostCategory  PostCategory
	PostedBy      User
}
