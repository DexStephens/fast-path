package graph

import "net/http"

type BaseHandler struct {}

func (h *BaseHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch {
	case req.Method == http.MethodPost:
		h.CreateGraph(res, req)
		return
	default:
		return
	}
}

func (h *BaseHandler) CreateGraph(res http.ResponseWriter, req *http.Request) {}