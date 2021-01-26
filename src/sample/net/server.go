package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", hello)
	log.Println("Starting v1 server ...")
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func hello(c Context) func(http.ResponseWriter, http.Request) {
	return func(c.W, c.r) {

	}
}

type Context struct {
	W http.ResponseWriter
	r *http.Request
}
