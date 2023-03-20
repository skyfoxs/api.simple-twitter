package model_test

import (
	"testing"

	"github.com/skyfoxs/api.simple-twitter/pkg/encrypt"
	"github.com/skyfoxs/api.simple-twitter/pkg/idata"
	"github.com/skyfoxs/api.simple-twitter/pkg/model"
)

func TestGetByID_ShouldReturnPop(t *testing.T) {
	um := model.NewUserModel()
	um.Add(pop)

	e := um.GetByID(pop.ID)

	if e == nil {
		t.Error("expect result to not be nil")
		return
	}
	if e.Email != pop.Email {
		t.Errorf("expect email to be %v but got %v", pop.Email, e.Email)
	}
	if e.Firstname != pop.Firstname {
		t.Errorf("expect email to be %v but got %v", pop.Firstname, e.Firstname)
	}
}

func TestGetByID_ShouldReturnJohn(t *testing.T) {
	um := model.NewUserModel()
	um.Add(pop)
	um.Add(john)

	e := um.GetByID(john.ID)

	if e == nil {
		t.Error("expect result to not be nil")
		return
	}
	if e.Email != john.Email {
		t.Errorf("expect email to be %v but got %v", john.Email, e.Email)
	}
	if e.Firstname != john.Firstname {
		t.Errorf("expect email to be %v but got %v", john.Firstname, e.Firstname)
	}
}

func TestGetByID_ShouldReturnNil(t *testing.T) {
	um := model.NewUserModel()
	um.Add(pop)
	um.Add(john)

	e := um.GetByID("random-id")

	if e != nil {
		t.Error("expect result to be nil")
	}
}

func TestGetByEmail_ShouldReturnPop(t *testing.T) {
	um := model.NewUserModel()
	um.Add(pop)

	e := um.GetByEmail(pop.Email)

	if e == nil {
		t.Error("expect result to not be nil")
		return
	}
	if e.Email != pop.Email {
		t.Errorf("expect email to be %v but got %v", pop.Email, e.Email)
	}
	if e.Firstname != pop.Firstname {
		t.Errorf("expect email to be %v but got %v", pop.Firstname, e.Firstname)
	}
}

func TestGetByEmail_ShouldReturnJohn(t *testing.T) {
	um := model.NewUserModel()
	um.Add(pop)
	um.Add(john)

	e := um.GetByEmail(john.Email)

	if e == nil {
		t.Error("expect result to not be nil")
		return
	}
	if e.Email != john.Email {
		t.Errorf("expect email to be %v but got %v", john.Email, e.Email)
	}
	if e.Firstname != john.Firstname {
		t.Errorf("expect email to be %v but got %v", john.Firstname, e.Firstname)
	}
}

func TestGetByEmail_ShouldReturnNil(t *testing.T) {
	um := model.NewUserModel()
	um.Add(pop)
	um.Add(john)

	e := um.GetByEmail("random-email")

	if e != nil {
		t.Error("expect result to be nil")
	}
}

func TestGetByEmailAndPassword_ShouldReturnPop(t *testing.T) {
	um := model.NewUserModel()
	um.Add(pop)
	um.Add(john)

	if e, valid := um.Get(pop.Email, "password"); valid {
		if e == nil {
			t.Error("expect result to not be nil")
			return
		}
		if e.Email != pop.Email {
			t.Errorf("expect email to be %v but got %v", pop.Email, e.Email)
		}
		if e.Firstname != pop.Firstname {
			t.Errorf("expect email to be %v but got %v", pop.Firstname, e.Firstname)
		}
	} else {
		t.Error("expect to be valid password")
	}
}

func TestGetByEmailAndPassword_ShouldReturnNil(t *testing.T) {
	um := model.NewUserModel()
	um.Add(pop)
	um.Add(john)

	if e, valid := um.Get(pop.Email, "wrong-password"); valid {
		t.Error("expect password to not valid")
		if e != nil {
			t.Error("expect result to be nil")
			return
		}
	}
}

var pop = idata.Profile{
	ID:        "pop-uuid",
	Firstname: "Pakornpat",
	Lastname:  "Sinjiranon",
	Email:     "skyfox.ku@gmail.com",
	Password:  encrypt.MD5str("password"),
}

var john = idata.Profile{
	ID:        "john-uuid",
	Firstname: "John",
	Lastname:  "Doe",
	Email:     "john@gmail.com",
	Password:  encrypt.MD5str("johnpass"),
}
