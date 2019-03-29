package main

import (
	"fmt"
	"lib/server/http"
	http2 "net/http"
)

func main() {

	fmt.Println("Hello Tyrion!")

	app := http.New()

	app.Init(&http.Options{
		IP:   "",
		Port: 3210,
	})

	http2.HandleFunc("/demo", func(writer http2.ResponseWriter, request *http2.Request) {
		fmt.Fprintf(writer, "Hello")
	})

	if err := app.Run(); err != nil {
		panic(err.Error())
	}
}
