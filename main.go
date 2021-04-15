package main

import (
	"golangweb-basic/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/vento", handler.VentoHandler)
	mux.HandleFunc("/product", handler.ProductHandler)

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}
