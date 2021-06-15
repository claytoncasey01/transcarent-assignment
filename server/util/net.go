package util

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Fetch(url string, ch chan<- []byte, c echo.Context) {
	resp, err := http.Get(url)
	if err != nil {
		c.Logger().Fatalf(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Logger().Fatalf(err.Error())
	}

	ch <- body

}
