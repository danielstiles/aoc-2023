package hikes

type Edge struct {
	Start Pos
	End   Pos
	Len   int
}

type Graph struct {
	Nodes []Pos
	Edges []Edge
}

func NewGraph(m *Maze, nodes []Pos, slippery bool) *Graph {
	moves := make([]Pos, 4)
	var edges []Edge
	for row, rowArr := range m.Tiles {
		for col, tile := range rowArr {
			if tile == Wall {
				continue
			}
			node := Pos{Row: row, Col: col}
			m.ValidMoves(node, slippery, moves)
			count := 0
			for _, pos := range moves {
				if pos.Steps == -1 {
					continue
				}
				count += 1
			}
			if count > 2 {
				nodes = append(nodes, node)
			}
		}
	}
	for _, node := range nodes {
		m.ValidMoves(node, slippery, moves)
		for _, pos := range moves {
			if pos.Steps == -1 {
				continue
			}
			for _, other := range nodes {
				if pos.Row == other.Row && pos.Col == other.Col {
					edges = append(edges, Edge{Start: node, End: other, Len: pos.Steps})
				}
			}
		}
	}
	return &Graph{
		Nodes: nodes,
		Edges: edges,
	}
}

func (g *Graph) GetEdges(node Pos) []Edge {
	var edges []Edge
	for _, edge := range g.Edges {
		if edge.Start.Row == node.Row && edge.Start.Col == node.Col {
			edges = append(edges, edge)
		}
	}
	return edges
}
