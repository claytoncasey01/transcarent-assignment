package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/claytoncasey01/transcarent-assignment/router"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserPosts(t *testing.T) {
	e := router.New()
	var h *Handler = NewHandler("https://jsonplaceholder.typicode.com/users", "https://jsonplaceholder.typicode.com/posts")

	req := httptest.NewRequest(echo.GET, "http://localhost:8080/v1/user-posts/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/user-posts/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	assert.NoError(t, h.GetUserPosts(c))
	res := userPostsResponse{}
	if assert.Equal(t, http.StatusOK, rec.Code) {
		json.Unmarshal(rec.Body.Bytes(), &res)
		assert.Equal(t, "Leanne Graham", res.UserInfo.Name)
		assert.Equal(t, "Bret", res.UserInfo.Username)
		assert.Equal(t, "Sincere@april.biz", res.UserInfo.Email)
		// Assert first and last posts in the array
		assert.Equal(t, 1, res.Posts[0].UserId)
		assert.Equal(t, 1, res.Posts[0].Id)
		assert.Equal(t, "sunt aut facere repellat provident occaecati excepturi optio reprehenderit", res.Posts[0].Title)
		assert.Equal(t, "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto", res.Posts[0].Body)
		assert.Equal(t, 1, res.Posts[9].UserId)
		assert.Equal(t, 10, res.Posts[9].Id)
		assert.Equal(t, "optio molestias id quia eum", res.Posts[9].Title)
		assert.Equal(t, "quo et expedita modi cum officia vel magni\ndoloribus qui repudiandae\nvero nisi sit\nquos veniam quod sed accusamus veritatis error", res.Posts[9].Body)
	}
}
