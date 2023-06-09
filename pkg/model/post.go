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

func (m *PostModel) AddComment(id string, c idata.Post) (ok bool) {
	p := m.GetByID(id)
	if p == nil {
		return false
	}
	if p.Comments == nil {
		p.Comments = []idata.Post{c}
	} else {
		p.Comments = append(p.Comments, c)
	}
	return true
}

func (m *PostModel) GetByID(id string) *idata.Post {
	for i, v := range m.posts {
		if v.ID == id {
			return &m.posts[i]
		}
	}
	return nil
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
