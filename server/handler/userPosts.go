package handler

import (
	"net/http"
	"strconv"

	"github.com/claytoncasey01/transcarent-assignment/model"
	"github.com/labstack/echo/v4"
)

type userPostsResponse struct {
	Id       int `json:"id"`
	UserInfo struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"userInfo"`
	Posts []*model.Post `json:"posts"`
}

func newUserPostsResponse(user *model.User, posts []*model.Post) userPostsResponse {
	userPostsResponse := userPostsResponse{
		Id:    user.Id,
		Posts: posts,
	}
	userPostsResponse.UserInfo.Name = user.Name
	userPostsResponse.UserInfo.Username = user.Username
	userPostsResponse.UserInfo.Email = user.Email

	return userPostsResponse
}

func (h *Handler) GetUserPosts(c echo.Context) error {
	// Get the userId from the request
	userId := c.Param("id")
	parsedUserId, err := strconv.Atoi(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "UserId must be an integer")
	}

	// Clamp the value to valid userIds, probably not necessary but might
	// as well handle it since we know the bounds
	if parsedUserId < 1 || parsedUserId > 10 {
		return c.JSON(http.StatusBadRequest, "Invalid userId supplied, must be between 1 and 10")
	}

	// Get the user
	userResp, err := http.Get(h.usersResource + "/" + userId)
	if err != nil {
		c.Logger().Fatalf(err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}
	user, err := model.NewUserFromResponse(userResp)
	if err != nil {
		c.Logger().Fatalf(err.Error())
		return c.JSON(http.StatusInternalServerError, "Unable to process your request at this time")
	}

	// Get the posts
	postsResp, err := http.Get(h.postsResource + "?userId=" + userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	posts, err := model.NewPostsSliceFromResponse(postsResp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, newUserPostsResponse(user, posts))
}

func (h *Handler) Test(c echo.Context) error {
	return c.String(http.StatusOK, "{}")
}
