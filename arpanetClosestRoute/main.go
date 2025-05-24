package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func normalizeName(input string) string {
	return strings.Title(strings.ToLower(input))
}

func main() {
	// Create undirected graph from directed graph
	undirectedGraph := makeUndirected(graph)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter start point: ")
	startInput, _ := reader.ReadString('\n')
	start := normalizeName(strings.TrimSpace(startInput))

	fmt.Print("Enter end point: ")
	endInput, _ := reader.ReadString('\n')
	end := normalizeName(strings.TrimSpace(endInput))

	path, dist, err := Dijkstra(undirectedGraph, start, end)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Shortest path from %s to %s is: ", start, end)
	for i, node := range path {
		fmt.Print(node)
		if i < len(path)-1 {
			fmt.Print(" -> ")
		}
	}
	fmt.Printf("\nTotal distance: %d\n", dist)
}
