package hikes

const (
	Empty = byte('.')
	Wall  = byte('#')
	Down  = byte('v')
	Up    = byte('^')
	Right = byte('>')
	Left  = byte('<')
)

type Maze struct {
	Tiles [][]byte
}

func (m *Maze) AddLine(line []byte) {
	m.Tiles = append(m.Tiles, line)
}

type Pos struct {
	Row   int
	Col   int
	Steps int
}

func (m *Maze) ValidMoves(pos Pos, slippery bool, moves []Pos) {
	moves[0].Steps = -1
	moves[1].Steps = -1
	moves[2].Steps = -1
	moves[3].Steps = -1
	row := pos.Row
	col := pos.Col
	steps := pos.Steps + 1
	if row < 0 || row > len(m.Tiles) || col < 0 || col > len(m.Tiles[0]) {
		return
	}
	if row > 0 && m.Tiles[row-1][col] != Wall {
		newPos := Pos{Row: row - 1, Col: col, Steps: steps}
		newPos = m.follow(newPos, pos, slippery)
		moves[0] = newPos
	}
	if row < len(m.Tiles)-1 && m.Tiles[row+1][col] != Wall {
		newPos := Pos{Row: row + 1, Col: col, Steps: steps}
		newPos = m.follow(newPos, pos, slippery)
		moves[1] = newPos
	}
	if col > 0 && m.Tiles[row][col-1] != Wall {
		newPos := Pos{Row: row, Col: col - 1, Steps: steps}
		newPos = m.follow(newPos, pos, slippery)
		moves[2] = newPos
	}
	if col < len(m.Tiles[0])-1 && m.Tiles[row][col+1] != Wall {
		newPos := Pos{Row: row, Col: col + 1, Steps: steps}
		newPos = m.follow(newPos, pos, slippery)
		moves[3] = newPos
	}
}

func (m *Maze) follow(pos, prev Pos, slippery bool) Pos {
	for m.CheckBounds(pos) && m.Tiles[pos.Row][pos.Col] != Wall {
		next := pos
		next.Steps += 1
		row := pos.Row
		col := pos.Col
		count := 0
		if slippery {
			switch m.Tiles[row][col] {
			case Down:
				count += 1
				next.Row += 1
			case Up:
				count += 1
				next.Row -= 1
			case Right:
				count += 1
				next.Col += 1
			case Left:
				count += 1
				next.Col -= 1
			}
		}
		if !slippery || m.Tiles[row][col] == Empty {
			if row > 0 && prev.Row != row-1 && m.Tiles[row-1][col] != Wall {
				count += 1
				next.Row -= 1
			}
			if row < len(m.Tiles)-1 && prev.Row != row+1 && m.Tiles[row+1][col] != Wall {
				count += 1
				next.Row += 1
			}
			if col > 0 && prev.Col != col-1 && m.Tiles[row][col-1] != Wall {
				count += 1
				next.Col -= 1
			}
			if col < len(m.Tiles[0])-1 && prev.Col != col+1 && m.Tiles[row][col+1] != Wall {
				count += 1
				next.Col += 1
			}
		}
		if count != 1 {
			return pos
		}
		prev = pos
		pos = next
	}
	return pos
}

func (m *Maze) CheckBounds(pos Pos) bool {
	return !(pos.Row < 0 || pos.Row >= len(m.Tiles) || pos.Col < 0 || pos.Col >= len(m.Tiles[0]))
}

func (m *Maze) FindLongestPath(start, end Pos, slippery bool) int {
	graph := NewGraph(m, []Pos{start, end}, slippery)
	var paths []*Path
	var finishedPaths []*Path
	for _, edge := range graph.GetEdges(start) {
		newPath := NewPath(m, start)
		next := edge.End
		next.Steps = edge.Len
		newPath.Move(next)
		paths = append(paths, newPath)
	}
	var first Pos
	for len(paths) > 0 {
		curr := paths[len(paths)-1]
		if curr.Contains(end) {
			finishedPaths = append(finishedPaths, curr)
			paths = paths[:len(paths)-1]
			continue
		}
		edges := graph.GetEdges(curr.Pos)
		first.Steps = -1
		for _, edge := range edges {
			if !curr.Contains(edge.End) {
				if first.Steps == -1 {
					first = edge.End
					first.Steps = curr.Pos.Steps + edge.Len
					continue
				}
				newPath := curr.Copy()
				next := edge.End
				next.Steps = curr.Pos.Steps + edge.Len
				newPath.Move(next)
				paths = append(paths, newPath)
			}
		}
		if first.Steps == -1 {
			paths = paths[:len(paths)-1]
			continue
		}
		curr.Move(first)
	}
	var max int
	for _, path := range finishedPaths {
		if path.Length(end) > max {
			max = path.Length(end)
		}
	}
	return max
}
