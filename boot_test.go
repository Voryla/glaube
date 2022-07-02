package glaube

import (
	"github.com/Voryla/Glaube/berouter"
	"net/http"
	"testing"
)

func TestName(t *testing.T) {
	r := berouter.Default()
	r.Get("/hello", func(c *berouter.Context) {
		a := berouter.H{"message": "hello Glaube!"}
		c.JSON(http.StatusOK, a)
	})
	r.Run(":8080")
}
