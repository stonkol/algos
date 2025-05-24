// dijkstra.go
package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Item for priority queue
type Item struct {
	node     string
	priority int
	index    int
}

// PriorityQueue implements heap.Interface
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func Dijkstra(graph map[string][]Edge, start, end string) ([]string, int, error) {
	dist := make(map[string]int)
	prev := make(map[string]string)
	for node := range graph {
		dist[node] = math.MaxInt32
	}
	if _, ok := graph[start]; !ok {
		return nil, 0, fmt.Errorf("start node %q not found", start)
	}
	if _, ok := graph[end]; !ok {
		return nil, 0, fmt.Errorf("end node %q not found", end)
	}
	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: start, priority: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		if current.node == end {
			break
		}
		for _, edge := range graph[current.node] {
			alt := dist[current.node] + edge.Weight
			if alt < dist[edge.To] {
				dist[edge.To] = alt
				prev[edge.To] = current.node
				heap.Push(pq, &Item{node: edge.To, priority: alt})
			}
		}
	}

	if dist[end] == math.MaxInt32 {
		return nil, 0, fmt.Errorf("no path found from %s to %s", start, end)
	}

	// Reconstruct path
	path := []string{}
	for at := end; at != start; at = prev[at] {
		path = append([]string{at}, path...)
	}
	path = append([]string{start}, path...)

	return path, dist[end], nil
}
