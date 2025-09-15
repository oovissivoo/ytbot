package main

import (
	"fmt"
	"net/http"
)

func App() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Hello world!")
	})
	return router
}

func main() {
	app := App()
	server := http.Server{
		Addr:    ":8080",
		Handler: app,
	}
	fmt.Println("YTbot listening")
	server.ListenAndServe()
}
