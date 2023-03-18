package model

import (
	"github.com/skyfoxs/api.simple-twitter/pkg/encrypt"
	"github.com/skyfoxs/api.simple-twitter/pkg/idata"
)

type UserModel struct {
	users []idata.Profile
}

func NewUser() *UserModel {
	return &UserModel{
		users: []idata.Profile{},
	}
}

func (a *UserModel) Get(e string, p string) (up *idata.Profile, valid bool) {
	up = a.GetByEmail(e)
	return up, up != nil && up.Password == encrypt.MD5str(p)
}

func (a *UserModel) GetByID(id string) *idata.Profile {
	for i, v := range a.users {
		if v.ID == id {
			return &a.users[i]
		}
	}
	return nil
}

func (a *UserModel) GetByEmail(e string) *idata.Profile {
	for i, v := range a.users {
		if v.Email == e {
			return &a.users[i]
		}
	}
	return nil
}

func (a *UserModel) Add(p idata.Profile) {
	a.users = append(a.users, p)
}
