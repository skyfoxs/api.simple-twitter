package model_test

import (
	"testing"

	"github.com/skyfoxs/api.simple-twitter/pkg/encrypt"
	"github.com/skyfoxs/api.simple-twitter/pkg/idata"
	"github.com/skyfoxs/api.simple-twitter/pkg/model"
)

func TestGetByID_ShouldReturnPop_WhenUsingPopID(t *testing.T) {
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

func TestGetByID_ShouldReturnJohn_WhenUsingJohnID(t *testing.T) {
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

func TestGetByID_ShouldReturnNil_WhenUserIDNotExist(t *testing.T) {
	um := model.NewUserModel()
	um.Add(pop)
	um.Add(john)

	e := um.GetByID("random-id")

	if e != nil {
		t.Error("expect result to be nil")
	}
}

func TestGetByEmail_ShouldReturnPop_WhenUsingPopEmail(t *testing.T) {
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

func TestGetByEmail_ShouldReturnJohn_WhenUsingJohnEmail(t *testing.T) {
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

func TestGetByEmail_ShouldReturnNil_WhenEmailNotExist(t *testing.T) {
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

	if e, valid := um.Get(pop.Email, "pop-password"); valid {
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

func TestGetByEmailAndPassword_ShouldReturnNil_WhenPasswordIsWrong(t *testing.T) {
	um := model.NewUserModel()
	um.Add(pop)
	um.Add(john)

	if e, valid := um.Get(pop.Email, "john-password"); valid {
		t.Error("expect password to not valid")
		if e != nil {
			t.Error("expect result to be nil")
			return
		}
	}
}

func TestAddFollowing_ShouldNotSuccess_WhenUserIDNotExist(t *testing.T) {
	um := model.NewUserModel()
	um.Add(pop)
	um.Add(john)
	if ok := um.AddFollowing(pop.ID, "not-exist-id"); ok {
		t.Error("expect can not add non existed user into pop following list")
		return
	}
}

func TestAddFollowing_ShouldNotSuccess_WhenFollowingYourself(t *testing.T) {
	um := model.NewUserModel()
	um.Add(pop)
	um.Add(john)
	if ok := um.AddFollowing(pop.ID, pop.ID); ok {
		t.Error("expect can not add yourself into following list")
		return
	}
}

func TestGetFollowing(t *testing.T) {
	um := model.NewUserModel()
	um.Add(pop)
	um.Add(john)
	if ok := um.AddFollowing(pop.ID, john.ID); !ok {
		t.Error("expect can add john into pop following list")
		return
	}
	r := um.GetFollowing(pop.ID)
	if len(r) != 1 {
		t.Error("expect to have 1 user in pop following list")
		return
	}
	if r[0].ID != john.ID {
		t.Error("expect john in pop following list")
		return
	}
}

func TestDeleteFollowing(t *testing.T) {
	um := model.NewUserModel()
	um.Add(pop)
	um.Add(john)
	if ok := um.AddFollowing(pop.ID, john.ID); !ok {
		t.Error("expect can add john into pop following list")
		return
	}
	if ok := um.DeleteFollowing(pop.ID, john.ID); !ok {
		t.Error("expect can delete john from pop following list")
		return
	}
	r := um.GetFollowing(pop.ID)
	if len(r) != 0 {
		t.Error("expect pop's following list to be empty")
		return
	}
}

var pop = idata.Profile{
	ID:        "pop-uuid",
	Firstname: "Pakornpat",
	Lastname:  "Sinjiranon",
	Email:     "skyfox.ku@gmail.com",
	Password:  encrypt.MD5str("pop-password"),
}

var john = idata.Profile{
	ID:        "john-uuid",
	Firstname: "John",
	Lastname:  "Doe",
	Email:     "john@gmail.com",
	Password:  encrypt.MD5str("john-password"),
}

var alex = idata.Profile{
	ID:        "alex-uuid",
	Firstname: "Alex",
	Lastname:  "Doe",
	Email:     "alex@gmail.com",
	Password:  encrypt.MD5str("alex-password"),
}
