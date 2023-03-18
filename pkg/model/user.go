package model

import (
	"github.com/skyfoxs/api.simple-twitter/pkg/encrypt"
	"github.com/skyfoxs/api.simple-twitter/pkg/idata"
)

type UserModel struct {
	users []idata.Profile
}

func NewUserModel() *UserModel {
	return &UserModel{
		users: []idata.Profile{},
	}
}

func (m *UserModel) Get(e string, p string) (up *idata.Profile, valid bool) {
	up = m.GetByEmail(e)
	return up, up != nil && up.Password == encrypt.MD5str(p)
}

func (m *UserModel) GetByID(id string) *idata.Profile {
	for i, v := range m.users {
		if v.ID == id {
			return &m.users[i]
		}
	}
	return nil
}

func (m *UserModel) GetByEmail(e string) *idata.Profile {
	for i, v := range m.users {
		if v.Email == e {
			return &m.users[i]
		}
	}
	return nil
}

func (m *UserModel) Add(p idata.Profile) {
	m.users = append(m.users, p)
}
