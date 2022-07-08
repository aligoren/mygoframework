package mygoframework

import (
	"net/http"

	"github.com/uptrace/bunrouter"
)

func (app *App) Get(path string, callback RequestHandlerFunc) error {
	app.router.GET(path, func(w http.ResponseWriter, req bunrouter.Request) error {

		app.ctx.response.ResponseWriter = w
		app.ctx.request.Request = req
		callback(app.ctx)

		return nil
	})

	return nil
}
