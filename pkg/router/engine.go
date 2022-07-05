package router

import (
	"github.com/Voryla/Glaube/pkg/common/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type H map[string]any

type Engine struct {
	router *httprouter.Router
	RouteGroup
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.router.ServeHTTP(w, r)
}

func (e *Engine) doHandler(ctx *Context, handler HandlerFunc) {
}

func (e *Engine) Run(addr string) {
	utils.PrintLogo()
	if err := http.ListenAndServe(addr, e); err != nil {

	}
}
func Default() *Engine {

	engine := &Engine{
		router: httprouter.New(),
		RouteGroup: RouteGroup{
			handlerChain: nil,
		},
	}
	engine.RouteGroup.Engine = engine
	return engine
}
