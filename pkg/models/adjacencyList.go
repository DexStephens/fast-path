package models

type AdjacencyList struct {
	Nodes map[int][]int `json:"nodes"`
}