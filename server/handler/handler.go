package handler

type Handler struct {
	usersResource string
	postsResource string
}

func NewHandler(usersResource string, postsResource string) *Handler {
	return &Handler{
		usersResource: usersResource,
		postsResource: postsResource,
	}
}
