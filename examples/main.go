package main

import (
	"log"
	"mygoframework"
	"net/http"
)

func main() {
	app, err := mygoframework.New(mygoframework.MyConfig{
		Name: "MyApp",
		Addr: ":8080",
	})

	if err != nil {
		log.Fatal(err)
	}

	app.Get("/", func(ctx *mygoframework.Context) error {
		return ctx.SendString("Hello world!")
	})

	app.Get("/user/:id", func(ctx *mygoframework.Context) error {

		return ctx.SetStatus(http.StatusNotFound).SendString("world 1")
	})

	app.Get("/json/:id", func(ctx *mygoframework.Context) error {

		id, _ := ctx.Param("id")

		return ctx.SetStatus(http.StatusBadGateway).SendJson(mygoframework.J{
			"id":   id,
			"name": "Ali",
		})
	})

	app.Get("/xml/:id", func(ctx *mygoframework.Context) error {

		id, _ := ctx.Param("id")

		type ExtraData struct {
			Name string
		}

		type User struct {
			ID   string
			Data ExtraData
		}

		user := User{
			ID: id,
			Data: ExtraData{
				Name: "Ali",
			},
		}

		return ctx.SendXml(user)
	})

	app.Start()
}
