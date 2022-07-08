## My Go Framework

Just a hobby. I tried something and it worked.

## Example

```go
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

	app.Get("/app/:id", func(ctx *mygoframework.Context) error {

		id, _ := ctx.Param("id")

		return ctx.SetStatus(http.StatusOK).SendJson(mygoframework.J{
			"id":   id,
			"name": "Ali",
		})
	})

	app.Start()
}
```