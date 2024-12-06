package util

import (
	"fmt"
	"slices"
)

type GraphList[T comparable] struct {
	adjList map[Vertex[T]][]Vertex[T]
}

type Vertex[T comparable] struct {
	Val T
}

func NewGraphMatrix[T comparable]() *GraphList[T] {
	return &GraphList[T]{
		adjList: make(map[Vertex[T]][]Vertex[T]),
	}
}

func (g *GraphList[T]) AddVertex(vertex Vertex[T]) {
	if _, exists := g.adjList[vertex]; !exists {
		g.adjList[vertex] = []Vertex[T]{}
	}
}

func (g *GraphList[T]) RemoveVertex(vertex Vertex[T]) {
	delete(g.adjList, vertex)
	for key, neighbors := range g.adjList {
		if slices.Contains(neighbors, vertex) {
			idx := slices.Index(neighbors, vertex)

			val := append(neighbors[:idx], neighbors[idx+1:]...)
			g.adjList[key] = val
		}
	}
}

func (g *GraphList[T]) AddEdge(from, to Vertex[T]) {
	fmt.Printf("Adding edge from %c to %c\n", from, to)
	g.AddVertex(from)
	g.AddVertex(to)

	g.adjList[from] = append(g.adjList[from], to)
	// g.adjList[to] = append(g.adjList[to], from) // undirected
}

func (g *GraphList[T]) Print() {
	fmt.Println("Adjancy List: ")
	for Vertex, neighbors := range g.adjList {
		fmt.Printf("%c -> %c\n", Vertex, neighbors)
	}
}

func (g *GraphList[T]) HasPath(start, end Vertex[T]) bool {
	path := g.BFS(start)
	return slices.Contains(path, end)
}

func (g *GraphList[T]) BFS(start Vertex[T]) []Vertex[T] {
	var visitedNodes []Vertex[T]
	visited := make(map[Vertex[T]]bool)

	queue := []Vertex[T]{start}

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		visitedNodes = append(visitedNodes, currentNode)
		for _, neighbor := range g.adjList[currentNode] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)

			}
		}

	}

	return visitedNodes
}

func (g *GraphList[T]) HasCycle(start Vertex[T]) bool {
	var visitedNodes []Vertex[T]
	visited := make(map[Vertex[T]]bool)

	queue := []Vertex[T]{start}

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		visitedNodes = append(visitedNodes, currentNode)
		for _, neighbor := range g.adjList[currentNode] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			} else {
				return true
			}
		}

	}

	return false
}

func (g *GraphList[T]) DFS(start Vertex[T]) {
	visited := make(map[Vertex[T]]bool)
	g.DFSUtil(start, visited)
}

func (g *GraphList[T]) DFSUtil(vertex Vertex[T], visited map[Vertex[T]]bool) {
	visited[vertex] = true
	fmt.Printf("%d ", vertex)

	for _, v := range g.adjList[vertex] {
		if !visited[v] {
			g.DFSUtil(v, visited)
		}
	}
}
