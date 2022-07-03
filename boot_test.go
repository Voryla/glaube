package glaube

import (
	"github.com/Voryla/Glaube/pkg/router"
	"net/http"
	"testing"
)

func TestName(t *testing.T) {
	r := router.Default()
	r.Get("/hello", func(c *router.Context) {
		a := router.H{"message": "hello glaube!"}
		c.JSON(http.StatusOK, a)
	})
	r.Run(":8080")
}
