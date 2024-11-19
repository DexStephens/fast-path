package route

import "net/http"

type RouteHandler struct {}

func (h *RouteHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch {
	default:
		return
	}
}