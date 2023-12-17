package factory

import (
	"slices"
	"strconv"
)

const (
	North = iota
	East
	South
	West
)

type City struct {
	Costs     [][]int
	MinLength int
	MaxLength int
}

func (c *City) AddLine(line []byte) {
	var row []int
	for _, c := range line {
		row = append(row, int(c)-48)
	}
	c.Costs = append(c.Costs, row)
}

type state struct {
	row     int
	col     int
	dir     int
	pathLen int
	cost    int
}

func (s state) hash() string {
	return strconv.Itoa(s.row) + "," + strconv.Itoa(s.col) + "," + strconv.Itoa(s.dir) + "," + strconv.Itoa(s.pathLen)
}

func (c *City) FindPath(startRow, startCol, endRow, endCol int) int {
	startState := state{
		row: startRow,
		col: startCol,
	}
	seen := make(map[string]int)
	next := c.getNextStates(startState, true)
	for _, s := range next {
		seen[s.hash()] = s.cost
	}
	slices.SortFunc(next, func(a, b state) int { return a.cost - b.cost })
	for next[0].row != endRow || next[0].col != endCol {
		newStates := c.getNextStates(next[0], false)
		for _, s := range newStates {
			if _, ok := seen[s.hash()]; ok {
				continue
			}
			seen[s.hash()] = s.cost
			next = append(next, s)
		}
		next = next[1:]
		slices.SortFunc(next, func(a, b state) int { return a.cost - b.cost })
	}
	return next[0].cost
}

func (c *City) getNextStates(curr state, start bool) []state {
	var possible []state
	if len(c.Costs) == 0 {
		return possible
	}
	if (start || (curr.dir == East || curr.dir == West) && curr.pathLen >= c.MinLength || (curr.dir == South && curr.pathLen < c.MaxLength)) && curr.row < len(c.Costs)-1 {
		pathLen := 1
		if curr.dir == South {
			pathLen = curr.pathLen + 1
		}
		possible = append(possible, state{
			row:     curr.row + 1,
			col:     curr.col,
			dir:     South,
			pathLen: pathLen,
			cost:    curr.cost + c.Costs[curr.row+1][curr.col],
		})
	}
	if (start || (curr.dir == East || curr.dir == West) && curr.pathLen >= c.MinLength || (curr.dir == North && curr.pathLen < c.MaxLength)) && curr.row > 0 {
		pathLen := 1
		if curr.dir == North {
			pathLen = curr.pathLen + 1
		}
		possible = append(possible, state{
			row:     curr.row - 1,
			col:     curr.col,
			dir:     North,
			pathLen: pathLen,
			cost:    curr.cost + c.Costs[curr.row-1][curr.col],
		})
	}
	if (start || (curr.dir == North || curr.dir == South) && curr.pathLen >= c.MinLength || (curr.dir == East && curr.pathLen < c.MaxLength)) && curr.col < len(c.Costs[0])-1 {
		pathLen := 1
		if curr.dir == East {
			pathLen = curr.pathLen + 1
		}
		possible = append(possible, state{
			row:     curr.row,
			col:     curr.col + 1,
			dir:     East,
			pathLen: pathLen,
			cost:    curr.cost + c.Costs[curr.row][curr.col+1],
		})
	}
	if (start || (curr.dir == North || curr.dir == South) && curr.pathLen >= c.MinLength || (curr.dir == West && curr.pathLen < c.MaxLength)) && curr.col > 0 {
		pathLen := 1
		if curr.dir == West {
			pathLen = curr.pathLen + 1
		}
		possible = append(possible, state{
			row:     curr.row,
			col:     curr.col - 1,
			dir:     West,
			pathLen: pathLen,
			cost:    curr.cost + c.Costs[curr.row][curr.col-1],
		})
	}
	return possible
}
