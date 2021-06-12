package router

import (
	"testing"

	"github.com/labstack/echo/v4"
)

// Make sure we get a new Echo.
// Probably unecessary but all about that coverage
func TestNew(t *testing.T) {
	var i interface{} = New()

	_, ok := i.(*echo.Echo)

	if !ok {
		t.Fatalf("Did not recieve the expected type (*echo.Echo)")
	}
}
