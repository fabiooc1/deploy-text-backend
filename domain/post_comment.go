package domain

import "time"

type PostComment struct {
	ID          int64
	Content     string
	CommentedAt time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
	PostID      int64
	CommentedBy User
	Parent      *PostComment
}
