package domain

type PostFavorite struct {
	ID     int64
	UserID int64
	Post   Post
}
