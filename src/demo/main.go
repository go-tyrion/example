package main

import (
	"demo/api"
	"fmt"
	"lib/core/error"
	"lib/server/http"
	"log"
)

func init() {
	fmt.Println("Hello Tyrion!")
}
func main() {
	runHttpServer()

	/*
		client := redis.NewClient(&redis.Options{
			Addr: "",
		})
		client.Get("here")

		redis.NewClusterClient(&redis.ClusterOptions{})
	*/

	e := err()
	if e != nil {
		fmt.Println("e:", e)
	}

	log.Print()
}

func err() *error.Error {
	code := error.ErrorCode(1)

	// return error.New("msg")
	return error.NewWithCode(code, "message")
}

func runHttpServer() {
	app := http.New()

	app.Init(&http.Options{
		Port:                3210,
		IgnorePathLastSlash: true,
	})

	app.AddLogic("/index", new(api.Index))

	app.Get("/", func(c *http.Context) {
		c.OkJSON(map[string]interface{}{
			"status":  0,
			"message": "success",
		})
	})

	app.Get("/demo", func(ctx *http.Context) {
		ctx.String(200, "Hello my framework.")
	})

	app.Get("/test", func(ctx *http.Context) {
		ctx.JSON(200, map[string]string{
			"status": "0",
			"msg":    "success",
			"demo":   ctx.Get("demo"),
			"foo":    ctx.Post("foo"),
		})
	})

	if err := app.Run(); err != nil {
		panic(err.Error())
	}

}
