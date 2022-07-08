package mygoframework

import "net/http"

type Response struct {
	http.ResponseWriter
}

func (ctx *Context) getStatusCode() int {

	if ctx.statusCode == 0 {
		ctx.statusCode = 200
	}

	return ctx.statusCode
}
