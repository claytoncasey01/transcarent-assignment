package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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

func NewPostsSliceFromResponse(postResonse *http.Response) ([]*Post, error) {
	posts := make([]*Post, 0)
	postBody, err := ioutil.ReadAll(postResonse.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(postBody, &posts)

	return posts, nil
}
