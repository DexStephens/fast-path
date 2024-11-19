package main

import (
	"fmt"
	"net/http"

	"github.com/DexStephens/fast-path/pkg/graph"
	"github.com/DexStephens/fast-path/pkg/route"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/graph", &graph.BaseHandler{})
	mux.Handle("/route", &route.RouteHandler{})

	http.ListenAndServe(":8080", mux)

	fmt.Println("Server started on port 8080")
}