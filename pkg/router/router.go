package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type IRoutes interface {
	Use(...HandlerFunc) IRoutes
	Handle(string, string, ...HandlerFunc) IRoutes
	Any(string, ...HandlerFunc) IRoutes
	GET(string, ...HandlerFunc) IRoutes
	POST(string, ...HandlerFunc) IRoutes
	DELETE(string, ...HandlerFunc) IRoutes
	PATCH(string, ...HandlerFunc) IRoutes
	PUT(string, ...HandlerFunc) IRoutes
	OPTIONS(string, ...HandlerFunc) IRoutes
	HEAD(string, ...HandlerFunc) IRoutes
	StaticFile(string, string) IRoutes
	Static(string, string) IRoutes
}
type IRouter interface {
	IRoutes
	Chain(handlerFunc ...HandlerFunc) *RouteGroup
}

type RouteGroup struct {
	Engine       *Engine
	handlerChain HandlerFuncChain
}

func (r *RouteGroup) Use(handlerFunc ...HandlerFunc) IRoutes {
	if r.handlerChain == nil {
		return r.Chain(handlerFunc...)
	}
	r.handlerChain = append(r.handlerChain, handlerFunc...)
	return r
}

func (r *RouteGroup) Handle(s string, s2 string, handlerFunc ...HandlerFunc) IRoutes {
	panic("implement me")
}

func (r *RouteGroup) Any(s string, handlerFunc ...HandlerFunc) IRoutes {
	panic("implement me")
}

func (r *RouteGroup) GET(path string, handlerFunc ...HandlerFunc) IRoutes {
	router := r.Chain(handlerFunc...)
	r.Engine.router.GET(path, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		ctx := &Context{
			W:      writer,
			R:      request,
			Params: params,
		}

		router.doHandler(ctx)
	})
	return router
}

func (r *RouteGroup) POST(s string, handlerFunc ...HandlerFunc) IRoutes {
	panic("implement me")
}

func (r *RouteGroup) DELETE(s string, handlerFunc ...HandlerFunc) IRoutes {
	panic("implement me")
}

func (r *RouteGroup) PATCH(s string, handlerFunc ...HandlerFunc) IRoutes {
	panic("implement me")
}

func (r *RouteGroup) PUT(s string, handlerFunc ...HandlerFunc) IRoutes {
	panic("implement me")
}

func (r *RouteGroup) OPTIONS(s string, handlerFunc ...HandlerFunc) IRoutes {
	panic("implement me")
}

func (r *RouteGroup) HEAD(s string, handlerFunc ...HandlerFunc) IRoutes {
	panic("implement me")
}

func (r *RouteGroup) StaticFile(s string, s2 string) IRoutes {
	panic("implement me")
}

func (r *RouteGroup) Static(s string, s2 string) IRoutes {
	panic("implement me")
}

func (r *RouteGroup) Chain(handlerFunc ...HandlerFunc) *RouteGroup {
	return &RouteGroup{
		Engine:       r.Engine,
		handlerChain: handlerFunc,
	}
}

func (r *RouteGroup) doHandler(ctx *Context) {
	for _, handler := range r.handlerChain {
		handler(ctx)
	}
}
