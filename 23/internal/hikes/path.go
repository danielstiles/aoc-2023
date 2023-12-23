package hikes

type Path struct {
	visited map[int]int
	maze    *Maze
	Pos     Pos
}

func NewPath(m *Maze, pos Pos) *Path {
	newPath := Path{
		visited: make(map[int]int),
		maze:    m,
	}
	newPath.Move(pos)
	return &newPath
}

func (p *Path) Copy() *Path {
	newVisited := make(map[int]int)
	for pos, steps := range p.visited {
		newVisited[pos] = steps
	}
	return &Path{visited: newVisited, maze: p.maze, Pos: p.Pos}
}

func (p *Path) Move(pos Pos) {
	if !p.maze.CheckBounds(pos) {
		return
	}
	p.visited[pos.Row*len(p.maze.Tiles[0])+pos.Col] = pos.Steps
	p.Pos = pos
}

func (p *Path) Contains(pos Pos) bool {
	return p.visited[pos.Row*len(p.maze.Tiles[0])+pos.Col] != 0
}

func (p *Path) Length(pos Pos) int {
	return p.visited[pos.Row*len(p.maze.Tiles[0])+pos.Col]
}
