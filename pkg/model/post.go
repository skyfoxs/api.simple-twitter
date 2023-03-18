package model

import "github.com/skyfoxs/api.simple-twitter/pkg/idata"

type PostModel struct {
	posts []idata.Post
}

func NewPostModel() *PostModel {
	return &PostModel{
		posts: []idata.Post{},
	}
}

func (m *PostModel) Add(p idata.Post) {
	m.posts = append(m.posts, p)
}
