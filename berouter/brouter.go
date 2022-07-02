package berouter

import (
	"github.com/fatih/color"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type H map[string]any

type Engine struct {
	router *httprouter.Router
}
type HandlerFunc func(*Context)

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	color.Blue("[Request]%s", r.URL.Path)
	e.router.ServeHTTP(w, r)
}

func (e *Engine) Get(path string, handler HandlerFunc) {
	e.router.GET(path, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		c := &Context{
			W: writer,
			R: request,
		}
		handler(c)
	})
}

func (e *Engine) Run(addr string) {
	printLogo()
	color.Blue("[Boot]welcome to use Glaube")
	if err := http.ListenAndServe(addr, e); err != nil {

	}
}
func Default() *Engine {
	return &Engine{router: httprouter.New()}
}
