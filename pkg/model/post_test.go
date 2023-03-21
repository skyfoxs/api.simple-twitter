package model_test

import (
	"testing"
	"time"

	"github.com/skyfoxs/api.simple-twitter/pkg/idata"
	"github.com/skyfoxs/api.simple-twitter/pkg/model"
)

func TestGetByUserID_ShouldReturnPostsFromPop_WhenUsingPopAsUserID(t *testing.T) {
	um := model.NewPostModel()
	um.Add(popPost1)
	um.Add(johnPost)

	e := um.GetByUserID(pop.ID)

	if len(e) != 1 {
		t.Error("expect result to have 1 post")
		return
	}
	if e[0].Message != popPost1.Message {
		t.Errorf("expect message to be %v but got %v", popPost1.Message, e[0].Message)
	}
}

func TestGetByUserID_ShouldReturnPostsFromJohn_WhenUsingJohnAsUserID(t *testing.T) {
	um := model.NewPostModel()
	um.Add(popPost1)
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
	um.Add(popPost1)
	um.Add(johnPost)

	e := um.GetByUserID("not-exist")

	if len(e) != 0 {
		t.Error("expect result to be empty list")
		return
	}
}

func TestGetByUserID_ShouldReturnSortedPosts(t *testing.T) {
	um := model.NewPostModel()
	um.Add(popPost1)
	um.Add(popPost2)

	e := um.GetByUserID(pop.ID)

	if len(e) != 2 {
		t.Error("expect result to have 2 posts")
		return
	}
	if e[0].Message != popPost2.Message {
		t.Errorf("expect message to be %v but got %v", popPost2.Message, e[0].Message)
	}
	if e[1].Message != popPost1.Message {
		t.Errorf("expect message to be %v but got %v", popPost1.Message, e[1].Message)
	}
}

func TestGetFeed_ShouldReturnSortedPostsFromUserIDCombinedWithPostsFromFollowingUserID(t *testing.T) {
	um := model.NewPostModel()
	um.Add(popPost1)
	um.Add(johnPost)
	um.Add(alexPost)
	um.Add(popPost2)

	e := um.GetFeed(pop.ID, []string{john.ID})

	if len(e) != 3 {
		t.Error("expect result to have 3 posts")
		return
	}
	if e[0].Message != popPost2.Message {
		t.Errorf("expect message to be %v but got %v", popPost2.Message, e[0].Message)
	}
	if e[1].Message != johnPost.Message {
		t.Errorf("expect message to be %v but got %v", johnPost.Message, e[1].Message)
	}
	if e[2].Message != popPost1.Message {
		t.Errorf("expect message to be %v but got %v", popPost1.Message, e[2].Message)
	}
}

var popPost1 = idata.Post{
	ID:       "pop-post-1",
	Message:  "Hello, I'm Pop",
	UserID:   pop.ID,
	Datetime: time.Date(2023, 3, 21, 10, 0, 0, 0, time.UTC),
}

var popPost2 = idata.Post{
	ID:       "pop-post-2",
	Message:  "Second Post!",
	UserID:   pop.ID,
	Datetime: time.Date(2023, 3, 21, 11, 0, 0, 0, time.UTC),
}

var johnPost = idata.Post{
	ID:       "john-post",
	Message:  "Hello, I'm John",
	UserID:   john.ID,
	Datetime: time.Date(2023, 3, 21, 10, 15, 0, 0, time.UTC),
}

var alexPost = idata.Post{
	ID:       "alex-post",
	Message:  "Hello, I'm Alex",
	UserID:   alex.ID,
	Datetime: time.Date(2023, 3, 21, 10, 20, 0, 0, time.UTC),
}
