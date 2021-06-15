package model

import (
	"encoding/json"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func NewPost(userId, id int, title, body string) *Post {
	return &Post{
		UserId: userId,
		Id:     id,
		Title:  title,
		Body:   body,
	}
}

func NewPostsSliceFromResponse(postResonse []byte) ([]*Post, error) {
	posts := make([]*Post, 0)
	json.Unmarshal(postResonse, &posts)

	return posts, nil
}
