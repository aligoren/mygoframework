package mygoframework

import (
	"errors"
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
)

type App struct {
	config MyConfig
	router *bunrouter.Router
	ctx    *Context

	Log *Logger
}

type Logger struct {
	log.Logger
}

type MyConfig struct {
	Name string
	Addr string

	log *Logger
}

func newLogger() *Logger {
	logger := &Logger{
		Logger: *log.Default(),
	}
	return logger
}

func New(config ...MyConfig) (*App, error) {

	app := &App{}

	if len(config) > 0 {
		app.config = config[0]
	}

	if app.config.Addr == "" {
		return nil, errors.New("addr must be specified")
	}

	if app.Log == nil {
		app.Log = newLogger()
	}

	app.config.log = app.Log

	app.ctx = &Context{}

	app.router = bunrouter.New()

	return app, nil
}

func (app *App) Start() {

	app.config.log.Printf("App started on %s\n", app.config.Addr)

	http.ListenAndServe(app.config.Addr, app.router)

}
