package model_test

import (
	"testing"

	"github.com/skyfoxs/api.simple-twitter/pkg/idata"
	"github.com/skyfoxs/api.simple-twitter/pkg/model"
)

func TestGetByUserID_ShouldReturnPostsFromPop_WhenUsingPopAsUserID(t *testing.T) {
	um := model.NewPostModel()
	um.Add(popPost)
	um.Add(johnPost)

	e := um.GetByUserID(pop.ID)

	if len(e) != 1 {
		t.Error("expect result to have 1 post")
		return
	}
	if e[0].Message != popPost.Message {
		t.Errorf("expect message to be %v but got %v", popPost.Message, e[0].Message)
	}
}

func TestGetByUserID_ShouldReturnPostsFromJohn_WhenUsingJohnAsUserID(t *testing.T) {
	um := model.NewPostModel()
	um.Add(popPost)
	um.Add(johnPost)

	e := um.GetByUserID(john.ID)

	if len(e) != 1 {
		t.Error("expect result to have 1 post")
		return
	}
	if e[0].Message != johnPost.Message {
		t.Errorf("expect message to be %v but got %v", johnPost.Message, e[0].Message)
	}
}

func TestGetByUserID_ShouldReturnEmptyPost_WhenNotHavePostWithGivenUserID(t *testing.T) {
	um := model.NewPostModel()
	um.Add(popPost)
	um.Add(johnPost)

	e := um.GetByUserID("not-exist")

	if len(e) != 0 {
		t.Error("expect result to be empty list")
		return
	}
}

var popPost = idata.Post{
	ID:      "pop-post",
	Message: "Hello, I'm Pop",
	UserID:  pop.ID,
}

var johnPost = idata.Post{
	ID:      "john-post",
	Message: "Hello, I'm John",
	UserID:  john.ID,
}
