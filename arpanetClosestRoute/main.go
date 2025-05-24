package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter start point: ")
	startInput, _ := reader.ReadString('\n')
	start := strings.TrimSpace(startInput)

	fmt.Print("Enter end point: ")
	endInput, _ := reader.ReadString('\n')
	end := strings.TrimSpace(endInput)

	path, dist, err := Dijkstra(graph, start, end)
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
