package main

import (
	"github.com/claytoncasey01/transcarent-assignment/handler"
	"github.com/claytoncasey01/transcarent-assignment/router"
)

func main() {
	r := router.New()

	v1 := r.Group("/v1")
	h := handler.NewHandler("https://jsonplaceholder.typicode.com/users", "https://jsonplaceholder.typicode.com/posts")
	h.Register(v1)

	r.Logger.Fatal(r.Start(":8080"))
}
