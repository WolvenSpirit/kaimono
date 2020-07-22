package main

import (
	"fmt"
	"net/http"
)

/*
* customRoutes provides a place to register custom functions.
* It is expected that you would rebuild the project after modifying this file.
 */
func customRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/example", func(wr http.ResponseWriter, r *http.Request) { fmt.Println("Hello world!") })
}
