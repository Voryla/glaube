package main

import (
	"github.com/Voryla/Glaube/pkg/router"
	"net/http"
)

func main() {
	r := router.Default()
	r.GET("/hello", func(context *router.Context) {
		a := router.H{"message": "hello glaube!"}
		context.JSON(http.StatusOK, a)
	})
	r.GET("/hello1", func(context *router.Context) {
		a := router.H{"message": "hello glaube!1"}
		context.JSON(http.StatusOK, a)
	}).Use(func(context *router.Context) {
		context.JSON(http.StatusOK, "USE")
	})
	r.Run(":8080")
}
