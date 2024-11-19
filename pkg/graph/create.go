package graph

import (
	"net/http"

	"github.com/DexStephens/fast-path/pkg/models"
)

type PostGraph struct {
	List models.AdjacencyList `json:"adjacency_list"`
	Matrix models.AdjacencyMatrix `json:"adjacency_matrix"`
}

func (h *BaseHandler) CreateGraph(res http.ResponseWriter, req *http.Request) {

}