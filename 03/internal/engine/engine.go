package engine

import (
	"regexp"
	"strconv"

	"golang.org/x/exp/slices"
	"golang.org/x/exp/slog"
)

var symbols = []byte{'%', '$', '@', '/', '#', '=', '+', '-', '*', '&'}
var numExpr = regexp.MustCompile("\\d+")

func GetIndices(line string) []int {
	var indices []int
	for i, c := range []byte(line) {
		if slices.Contains(symbols, c) {
			indices = append(indices, i-1, i, i+1)
		}
	}
	return indices
}

func GetPartTotal(line string, neighbors ...string) int {
	var total int
	var indices []int
	indices = getIndices(neighbors...)
	numbers := numExpr.FindAllStringIndex(line, -1)
	if numbers == nil {
		return 0
	}
	for _, loc := range numbers {
		for i := loc[0]; i < loc[1]; i++ {
			if slices.Contains(indices, i) {
				num, err := strconv.Atoi(line[loc[0]:loc[1]])
				if err != nil {
					slog.Error("Bad number", slog.Any("error", err))
				}
				total += num
				break
			}
		}
	}
	return total
}

func getIndices(from ...string) []int {
	var into []int
	for _, s := range from {
		indices := GetIndices(s)
		for _, i := range indices {
			into = append(into, i)
		}
	}
	return into
}
