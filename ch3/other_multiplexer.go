package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s! n", p.ByName("name"))
}

func bye(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "bye, %s! n", p.ByName("name"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	mux.GET("/bye/:name", bye)

	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
