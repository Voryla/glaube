package router

import (
	"encoding/json"
	"github.com/fatih/color"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (c *Context) JSON(code int, data any) {
	c.W.WriteHeader(code)
	d, _ := json.Marshal(data)
	c.W.Write(d)
	color.Green("[Response]%s", data)
}
