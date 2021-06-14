package handler

import "github.com/labstack/echo/v4"

func (h *Handler) Register(v1 *echo.Group) {
	userPosts := v1.Group("/user-posts")
	userPosts.GET("/:id", h.GetUserPosts)
}
