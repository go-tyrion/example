package main

import (
	"fmt"
	"lib/server/http"
	http2 "net/http"
)

func main() {

	fmt.Println("Hello Tyrion!")

	a := &http.Options{
		Port: 12354,
	}

	b := new(http.Options)
	*b = *a
	b.Port = 54321

	fmt.Print(a)
	fmt.Print(b)
}

func runHttpServer() {
	app := http.New()

	app.Init(&http.Options{
		IP:   "",
		Port: 3210,
	})

	app.InitByConfig("project.ini")

	http2.HandleFunc("/demo", func(writer http2.ResponseWriter, request *http2.Request) {
		fmt.Fprintf(writer, "Hello")
	})

	if err := app.Run(); err != nil {
		panic(err.Error())
	}

}
