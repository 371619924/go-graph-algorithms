package lgraph

type node uint

type edge struct {
	destination node
	label       rune
}

// LGraph is a function representing a directed labeled graph. If the node exists
// in the graph, the function returns true along with the set of outgoing edges
// from that node, otherwise false and nil.
type LGraph func(node) ([]edge, bool)

// FindSequence returns (S, true) if there is a sequence S of length k from node
// s to node t in graph g1 and S is not a sequence from s to t in graph g2; else
// it returns (nil, false).
func FindSequence(g1, g2 LGraph, s, t node, k uint) ([]rune, bool) {
	sequence, ok := findInGraph(g1, g2, s, t, k, []rune{}, s)
	if !ok {
		return nil, false
	}
	return sequence, true
}
func findInGraph(g1, g2 LGraph, current, target node, stepsLeft uint, path []rune, start node) ([]rune, bool) {
	_, exists := g1(current)
	if !exists {
		return nil, false
	}

	if stepsLeft == 0 {
		if current != target {
			return nil, false
		}
		if hasSequence(g2, start, target, path, 0) {
			return nil, false
		}

		result := make([]rune, len(path))
		copy(result, path)
		return result, true
	}

	edges, _ := g1(current)
	for _, e := range edges {
		path = append(path, e.label)

		result, ok := findInGraph(g1, g2, e.destination, target, stepsLeft-1, path, start)
		if ok {
			return result, true
		}

		path = path[:len(path)-1]
	}

	return nil, false
}

func hasSequence(g LGraph, current, target node, sequence []rune, index int) bool {
	edges, exists := g(current)
	if !exists {
		return false
	}

	if index == len(sequence) {
		return current == target
	}

	for _, e := range edges {
		if e.label == sequence[index] {
			if hasSequence(g, e.destination, target, sequence, index+1) {
				return true
			}
		}
	}

	return false
}
