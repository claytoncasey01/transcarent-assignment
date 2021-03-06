package handler

import (
	"net/http"
	"strconv"

	"github.com/claytoncasey01/transcarent-assignment/model"
	"github.com/claytoncasey01/transcarent-assignment/util"
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
		return echo.NewHTTPError(http.StatusBadRequest, "UserId must be an integer")
	}

	// Clamp the value to valid userIds, probably not necessary but might
	// as well handle it since we know the bounds
	if parsedUserId < 1 || parsedUserId > 10 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid userId supplied, must be between 1 and 10")
	}

	// Get the user
	// userResp, err := http.Get(h.usersResource + "/" + userId)
	userResp := make(chan []byte)
	postsResp := make(chan []byte)

	go util.Fetch(h.usersResource+"/"+userId, userResp, c)
	go util.Fetch(h.postsResource+"?userId="+userId, postsResp, c)

	user, err := model.NewUserFromResponse(<-userResp)
	if err != nil {
		c.Logger().Fatalf(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to process your request at this time")
	}

	// Get the posts
	posts, err := model.NewPostsSliceFromResponse(<-postsResp)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, newUserPostsResponse(user, posts))
}
