package util

import (
	"net/http/httptest"
	"testing"

	"github.com/claytoncasey01/transcarent-assignment/router"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	respChan := make(chan []byte)
	req := httptest.NewRequest(echo.GET, "http://localhost:8080/v1/user-posts/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := router.New().NewContext(req, rec)

	go Fetch("https://jsonplaceholder.typicode.com/users", respChan, c)
	resp := <-respChan
	assert.NotEmpty(t, resp)
	assert.GreaterOrEqual(t, len(resp), 0)
	assert.Equal(t, 5645, len(resp)) // Magic number but we know the length will never change

}
