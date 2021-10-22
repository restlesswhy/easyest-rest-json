package main

import (
	"net/http"

	"github.com/restlesswhy/new_test"
)


func main() {
	srv := newtest.NewServer()

	http.ListenAndServe(":8080", srv)
}