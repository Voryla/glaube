package router

import (
	"encoding/json"
	"github.com/fatih/color"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Context struct {
	W      http.ResponseWriter
	R      *http.Request
	Params httprouter.Params
}
type HandlerFunc func(*Context)

type HandlerFuncChain []HandlerFunc

func (ctx *Context) JSON(code int, data any) {
	ctx.W.WriteHeader(code)
	d, _ := json.Marshal(data)
	ctx.W.Write(d)
	color.Green("[Response]%s", data)
}
