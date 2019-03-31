package main

import (
	"fmt"
	"lib/server/http"
)

func init() {
	fmt.Println("Hello Tyrion!")
}
func main() {
	runHttpServer()
}

func runHttpServer() {
	app := http.New()

	app.Init(&http.Options{
		Port: 3210,
	})

	app.Get("/demo", func(ctx *http.Context) {
		ctx.String(200, "Hello my framework.")
	})

	if err := app.Run(); err != nil {
		panic(err.Error())
	}

}
