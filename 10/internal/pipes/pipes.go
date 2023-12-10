package pipes

import "golang.org/x/exp/slog"

const (
	up         = 1
	left       = 2
	down       = 4
	right      = 8
	directions = 15
	path       = 16
	interior   = 32
)

var connections = map[byte]int{
	'.': 0,
	'|': 5,
	'-': 10,
	'L': 9,
	'J': 3,
	'7': 6,
	'F': 12,
	'S': 15,
}

var start = byte('S')

type Area struct {
	Tiles [][]int
	Start []int
}

func (a *Area) AddRow(line []byte) {
	var row []int
	for i, c := range line {
		row = append(row, connections[c])
		if c == start {
			a.Start = []int{len(a.Tiles), i}
		}
	}
	a.Tiles = append(a.Tiles, row)
}

func (a *Area) CalcPath() int {
	var pathLen, firstDir, prevDir int
	currRow := a.Start[0]
	currCol := a.Start[1]
	if (a.Tiles[currRow+1][currCol] & up) != 0 {
		currRow += 1
		firstDir = down
		prevDir = up
	} else if (a.Tiles[currRow][currCol+1] & left) != 0 {
		currCol += 1
		firstDir = right
		prevDir = left
	} else {
		currRow -= 1
		firstDir = up
		prevDir = down
	}
	pathLen += 1
	for currRow != a.Start[0] || currCol != a.Start[1] {
		nextDir := a.Tiles[currRow][currCol] & (^prevDir) & directions
		a.addToPath(currRow, currCol, a.Tiles[currRow][currCol])
		switch nextDir {
		case up:
			currRow -= 1
			prevDir = down
		case left:
			currCol -= 1
			prevDir = right
		case down:
			currRow += 1
			prevDir = up
		case right:
			currCol += 1
			prevDir = left
		}
		pathLen += 1
	}
	a.addToPath(currRow, currCol, prevDir|firstDir)
	return pathLen
}

func (a *Area) GetInteriorCount() int {
	var count int
	for _, row := range a.Tiles {
		rowStr := ""
		for _, val := range row {
			if (val & path) == 0 {
				if (val & interior) != 0 {
					count += 1
					rowStr += "1"
				} else {
					rowStr += "0"
				}
			} else {
				rowStr += "P"
			}
		}
		slog.Info("map", slog.String("row", rowStr))
	}
	return count
}

func (a *Area) addToPath(row, col, value int) {
	a.Tiles[row][col] |= path
	if (value & right) != 0 {
		for i := row; i < len(a.Tiles); i++ {
			a.Tiles[i][col] ^= interior
		}
	}
}
