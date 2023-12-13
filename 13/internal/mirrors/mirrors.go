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

func FindReflection(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	for i := 1; i < len(grid[0]); i++ {
		if checkColReflection(grid, i) {
			return i * colMult
		}
	}
	for i := 1; i < len(grid); i++ {
		if checkRowReflection(grid, i) {
			return i * rowMult
		}
	}
	return 0
}

func compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func getCol(grid [][]int, i int) []int {
	var col []int
	for _, row := range grid {
		col = append(col, row[i])
	}
	return col
}

func checkColReflection(grid [][]int, col int) bool {
	for i := 0; i < col; i++ {
		if col+i >= len(grid[0]) {
			return true
		}
		if !compare(getCol(grid, col+i), getCol(grid, col-i-1)) {
			return false
		}
	}
	return true
}

func checkRowReflection(grid [][]int, row int) bool {
	for i := 0; i < row; i++ {
		if row+i >= len(grid) {
			return true
		}
		if !compare(grid[row+i], grid[row-i-1]) {
			return false
		}
	}
	return true
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
