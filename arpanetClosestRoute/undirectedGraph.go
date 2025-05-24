package main

func makeUndirected(g map[string][]Edge) map[string][]Edge {
	undirected := make(map[string][]Edge)
	for from, edges := range g {
		// Copy existing edges
		undirected[from] = append(undirected[from], edges...)
		// Add reverse edges
		for _, e := range edges {
			undirected[e.To] = append(undirected[e.To], Edge{To: from, Distance: e.Distance})
		}
	}
	return undirected
}
