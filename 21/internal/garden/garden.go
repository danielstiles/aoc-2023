package garden

import (
	"fmt"
)

const (
	empty = byte('.')
	rock  = byte('#')
	step  = byte('S')
)

type Garden struct {
	Tiles [][]byte
}

func (g *Garden) AddLine(line []byte) {
	g.Tiles = append(g.Tiles, line)
}

func (g *Garden) Iterate() int {
	var newGarden [][]byte
	var positions [][]int
	for i, line := range g.Tiles {
		var newLine []byte
		for j, tile := range line {
			switch tile {
			case step:
				newLine = append(newLine, empty)
				positions = append(positions, []int{i, j})
			default:
				newLine = append(newLine, tile)
			}
		}
		newGarden = append(newGarden, newLine)
	}
	steps := 0
	for _, pos := range positions {
		i := pos[0]
		j := pos[1]
		if i < len(newGarden)-1 && newGarden[i+1][j] == empty {
			newGarden[i+1][j] = step
			steps += 1
		}
		if i > 0 && newGarden[i-1][j] == empty {
			newGarden[i-1][j] = step
			steps += 1
		}
		if j < len(newGarden[i])-1 && newGarden[i][j+1] == empty {
			newGarden[i][j+1] = step
			steps += 1
		}
		if j > 0 && newGarden[i][j-1] == empty {
			newGarden[i][j-1] = step
			steps += 1
		}
	}
	g.Tiles = newGarden
	return steps
}

func (g *Garden) Print() {
	for _, line := range g.Tiles {
		out := ""
		for _, tile := range line {
			out += string(tile)
		}
		fmt.Print(out + "\n")
	}
}
