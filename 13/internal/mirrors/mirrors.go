package mirrors

import "log/slog"

const (
	rock = byte('#')
	dust = byte('.')
)

const (
	rockFlag = 1
	colMult  = 1
	rowMult  = 100
)

func ParseLine(line []byte) []int {
	var ret []int
	for _, c := range line {
		if c == rock {
			ret = append(ret, rockFlag)
		} else {
			ret = append(ret, 0)
		}
	}
	return ret
}

func FindReflection(grid [][]int, smudge bool) int {
	if len(grid) == 0 {
		return 0
	}
	for i := 1; i < len(grid[0]); i++ {
		differences := checkColReflection(grid, i)
		if !smudge && differences == 0 || smudge && differences == 1 {
			return i * colMult
		}
	}
	for i := 1; i < len(grid); i++ {
		differences := checkRowReflection(grid, i)
		if !smudge && differences == 0 || smudge && differences == 1 {
			return i * rowMult
		}
	}
	return 0
}

func compare(a, b []int) int {
	if len(a) != len(b) {
		return -1
	}
	var differences int
	for i := range a {
		if a[i] != b[i] {
			differences += 1
		}
	}
	return differences
}

func getCol(grid [][]int, i int) []int {
	var col []int
	for _, row := range grid {
		col = append(col, row[i])
	}
	return col
}

func checkColReflection(grid [][]int, col int) int {
	var differences int
	for i := 0; i < col; i++ {
		if col+i >= len(grid[0]) {
			return differences
		}
		differences += compare(getCol(grid, col+i), getCol(grid, col-i-1))
	}
	return differences
}

func checkRowReflection(grid [][]int, row int) int {
	var differences int
	for i := 0; i < row; i++ {
		if row+i >= len(grid) {
			return differences
		}
		differences += compare(grid[row+i], grid[row-i-1])
	}
	return differences
}

func PrintGrid(grid [][]int) {
	for _, row := range grid {
		rowStr := ""
		for _, val := range row {
			if val == rockFlag {
				rowStr += "#"
			} else {
				rowStr += "."
			}
		}
		slog.Info("grid", slog.String("row", rowStr))
	}
}
