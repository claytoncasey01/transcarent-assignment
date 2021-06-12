package main

import (
	"github.com/claytoncasey01/transcarent-assignment/router"
)

func main() {
	r := router.New()

	// v1 := r.Group("/api")

	r.Logger.Fatal(r.Start(":8080"))
}
