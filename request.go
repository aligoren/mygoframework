package mygoframework

import "github.com/uptrace/bunrouter"

type Request struct {
	bunrouter.Request
}

type RequestHandlerFunc func(ctx *Context) error
