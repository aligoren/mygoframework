package main

import (
	"log"
	"mygoframework"
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

		ctx.SetStatus(404)

		return ctx.SendString("world 1")
	})

	app.Get("/app/:id", func(ctx *mygoframework.Context) error {

		ctx.SetStatus(404)

		return ctx.SendJson(mygoframework.J{
			"name": "Ali",
		})
	})

	app.Start()
}
