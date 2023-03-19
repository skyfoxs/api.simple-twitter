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

func (m *PostModel) GetByUserID(uid string) []idata.Post {
	result := []idata.Post{}
	for _, v := range m.posts {
		if v.UserID == uid {
			result = append(result, v)
		}
	}
	return result
}
