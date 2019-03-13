package utils

import "github.com/valyala/fasthttp"

type Middleware func(h fasthttp.RequestHandler) fasthttp.RequestHandler
type MiddlewareFunc func(ctx *fasthttp.RequestCtx, args ...interface{})
type MiddlewareResultFunc func(ctx *fasthttp.RequestCtx, args ...interface{})

func NewMiddleware(middlewareFunc MiddlewareFunc, args ...interface{}) Middleware {
	return func(h fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			middlewareFunc(ctx, args)
			h(ctx)
		}
	}
}
