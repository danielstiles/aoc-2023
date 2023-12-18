package dig

import (
	"regexp"
	"strconv"
)

var lineExpr = regexp.MustCompile("(U|D|L|R) (\\d+) \\(#(\\w{6})\\)")

const (
	East  = 0
	South = 1
	West  = 2
	North = 3
)

type Command struct {
	Dir    int
	Length int
}

func ParseLine(line []byte, swap bool) Command {
	groups := lineExpr.FindSubmatch(line)
	var dir int
	switch groups[1][0] {
	case 'U':
		dir = North
	case 'D':
		dir = South
	case 'L':
		dir = West
	case 'R':
		dir = East
	}
	length, _ := strconv.Atoi(string(groups[2]))
	if swap {
		color, _ := strconv.ParseInt(string(groups[3]), 16, 32)
		dir = int(color) & 3
		length = int(color) >> 4
	}
	return Command{
		Dir:    dir,
		Length: length,
	}
}

func NewDigger() *Digger {
	site := &Polygon{}
	return &Digger{
		Site: site,
	}
}

func move(row, col int, command Command) (int, int) {
	switch command.Dir {
	case North:
		return row - command.Length, col
	case East:
		return row, col + command.Length
	case South:
		return row + command.Length, col
	case West:
		return row, col - command.Length
	}
	return -1, -1
}

const (
	path     = 1
	interior = 2
)

type Vector2 struct {
	X int
	Y int
}

func (v1 Vector2) Distance(v2 Vector2) int {
	diffX := v2.X - v1.X
	if diffX < 0 {
		diffX = -diffX
	}
	diffY := v2.Y - v1.Y
	if diffY < 0 {
		diffY = -diffY
	}
	return diffX + diffY
}

type Polygon struct {
	Vertices []Vector2
}

func (p *Polygon) AddVertex(point Vector2) {
	p.Vertices = append(p.Vertices, point)
}

func (p *Polygon) Area() int {
	if len(p.Vertices) < 3 {
		return 0
	}
	prev := p.Vertices[0]
	last := p.Vertices[len(p.Vertices)-1]
	total := last.X*prev.Y - last.Y*prev.X
	perim := last.Distance(prev)
	for _, v := range p.Vertices[1:] {
		total += prev.X*v.Y - prev.Y*v.X
		perim += prev.Distance(v)
		prev = v
	}
	if total < 0 {
		total = -total
	}
	return total/2 + perim/2 + 1
}

type Digger struct {
	Site *Polygon
	Row  int
	Col  int
}

func (d *Digger) Dig(command Command) {
	row, col := move(d.Row, d.Col, command)
	d.Site.AddVertex(Vector2{X: col, Y: row})
	d.Row = row
	d.Col = col
}
