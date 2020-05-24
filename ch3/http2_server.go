package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// go version < 1.6
	// http.ConfigureServer(&server, &http2.Server{})
	
	http.Handle("/hello", &handler)
	server.ListenAndServeTLS("cert.pem", "key.pem")		
}
