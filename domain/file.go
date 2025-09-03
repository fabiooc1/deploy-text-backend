package domain

import "time"

type File struct {
	ID           int64
	OriginalName string
	PathName     string
	MimeType     string
	Extension    string
	UploadedAt   *time.Time
	UploadedByID int64
}
