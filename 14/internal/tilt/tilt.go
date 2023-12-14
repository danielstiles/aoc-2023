package tilt

import "log/slog"

const (
	movingFlag     = 1
	stationaryFlag = 2
)

const (
	moving     = byte('O')
	stationary = byte('#')
	nothing    = byte('.')
)

const (
	North = iota
	East
	South
	West
)

type Grid struct {
	Grid [][]int
}

func (g *Grid) AddLine(line []byte) {
	var ret []int
	for _, c := range line {
		if c == moving {
			ret = append(ret, movingFlag)
		} else if c == stationary {
			ret = append(ret, stationaryFlag)
		} else {
			ret = append(ret, 0)
		}
	}
	g.Grid = append(g.Grid, ret)
}

func (g *Grid) Tilt(dir int) {
	if len(g.Grid) == 0 {
		return
	}
	switch dir {
	case North:
		for i := 0; i < len(g.Grid); i++ {
			for j := 0; j < len(g.Grid[0]); j++ {
				g.moveNorth(i, j)
			}
		}
	case East:
		for j := len(g.Grid[0]) - 1; j >= 0; j-- {
			for i := 0; i < len(g.Grid); i++ {
				g.moveEast(i, j)
			}
		}
	case South:
		for i := len(g.Grid) - 1; i >= 0; i-- {
			for j := 0; j < len(g.Grid[0]); j++ {
				g.moveSouth(i, j)
			}
		}
	case West:
		for j := 0; j < len(g.Grid); j++ {
			for i := 0; i < len(g.Grid); i++ {
				g.moveWest(i, j)
			}
		}
	}
}

func (g *Grid) moveNorth(row, col int) {
	if g.Grid[row][col] != movingFlag {
		return
	}
	for i := row - 1; i >= 0; i-- {
		if g.Grid[i][col] != 0 {
			break
		}
		g.Grid[i+1][col] = 0
		g.Grid[i][col] = movingFlag
	}
}

func (g *Grid) moveEast(row, col int) {
	if g.Grid[row][col] != movingFlag {
		return
	}
	for i := col + 1; i < len(g.Grid[0]); i++ {
		if g.Grid[row][i] != 0 {
			break
		}
		g.Grid[row][i-1] = 0
		g.Grid[row][i] = movingFlag
	}
}

func (g *Grid) moveSouth(row, col int) {
	if g.Grid[row][col] != movingFlag {
		return
	}
	for i := row + 1; i < len(g.Grid); i++ {
		if g.Grid[i][col] != 0 {
			break
		}
		g.Grid[i-1][col] = 0
		g.Grid[i][col] = movingFlag
	}
}

func (g *Grid) moveWest(row, col int) {
	if g.Grid[row][col] != movingFlag {
		return
	}
	for i := col - 1; i >= 0; i-- {
		if g.Grid[row][i] != 0 {
			break
		}
		g.Grid[row][i+1] = 0
		g.Grid[row][i] = movingFlag
	}
}

func (g *Grid) CalcWeight(dir int) int {
	total := 0
	for row, arr := range g.Grid {
		for col, val := range arr {
			if val == movingFlag {
				switch dir {
				case North:
					total += len(g.Grid) - row
				case East:
					total += len(g.Grid[0]) - col
				case South:
					total += row
				case West:
					total += col
				}
			}
		}
	}
	return total
}

func (g *Grid) Equal(other *Grid) bool {
	if len(g.Grid) != len(other.Grid) {
		return false
	}
	if len(g.Grid) == 0 && len(other.Grid) == 0 {
		return true
	}
	if len(g.Grid[0]) != len(other.Grid[0]) {
		return false
	}
	for row, arr := range g.Grid {
		for col, val := range arr {
			if val != other.Grid[row][col] {
				return false
			}
		}
	}
	return true
}

func (g *Grid) Copy(into *Grid) {
	var newGrid [][]int
	for _, arr := range g.Grid {
		var newRow []int
		for _, val := range arr {
			newRow = append(newRow, val)
		}
		newGrid = append(newGrid, newRow)
	}
	into.Grid = newGrid
}

func (g *Grid) Print() {
	for _, arr := range g.Grid {
		line := ""
		for _, val := range arr {
			if val == movingFlag {
				line += "O"
			} else if val == stationaryFlag {
				line += "#"
			} else {
				line += "."
			}
		}
		slog.Info("grid", slog.String("row", line))
	}
}
