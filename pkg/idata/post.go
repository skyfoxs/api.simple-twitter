package idata

import "time"

type Post struct {
	ID       string
	Message  string
	UserID   string
	Datetime time.Time
	Comments []Post
}
