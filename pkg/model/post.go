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
	m.posts = append([]idata.Post{p}, m.posts...)
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

func (m *PostModel) GetFeed(uid string, fid []string) []idata.Post {
	fidMap := map[string]bool{}
	result := []idata.Post{}
	fidMap[uid] = true
	for _, v := range fid {
		fidMap[v] = true
	}

	for _, v := range m.posts {
		if _, ok := fidMap[v.UserID]; ok {
			result = append(result, v)
		}
	}
	return result
}
