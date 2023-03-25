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

func (m *UserModel) GetFollowing(id string) []idata.Profile {
	p := m.GetByID(id)
	if p.Following == nil {
		return []idata.Profile{}
	}
	result := []idata.Profile{}
	for _, fid := range p.Following {
		result = append(result, *m.GetByID(fid))
	}
	return result
}

func (m *UserModel) AddFollowing(pid string, fid string) (ok bool) {
	p, _, ok := m.usersExistAndNotSame(pid, fid)
	if !ok {
		return false
	}

	if p.Following == nil {
		p.Following = []string{fid}
		return true
	}
	for _, v := range p.Following {
		if v == fid {
			return true
		}
	}
	p.Following = append(p.Following, fid)
	return true
}

func (m *UserModel) DeleteFollowing(pid string, fid string) (ok bool) {
	p, _, ok := m.usersExistAndNotSame(pid, fid)
	if !ok {
		return false
	}

	if p.Following == nil {
		p.Following = []string{}
		return true
	}
	r := []string{}
	for _, v := range p.Following {
		if v != fid {
			r = append(r, v)
		}
	}
	p.Following = r
	return true
}

func (m *UserModel) usersExistAndNotSame(pid, fid string) (p, f *idata.Profile, ok bool) {
	if pid == fid {
		return nil, nil, false
	}
	p = m.GetByID(pid)
	if p == nil {
		return nil, nil, false
	}
	f = m.GetByID(fid)
	if f == nil {
		return nil, nil, false
	}
	return p, f, true
}
