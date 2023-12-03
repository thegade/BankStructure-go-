package main

import (
	"bank/internal/di"
	"log"
	"net/http"
)

func main() {
	container := di.NewContainer()
	err := http.ListenAndServe("localhost:8080", container.HTTPRouter())
	if err != nil {
		log.Fatal(err)
	}
}
