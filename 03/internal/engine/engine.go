package engine

import (
	"regexp"
	"strconv"

	"golang.org/x/exp/slices"
	"golang.org/x/exp/slog"
)

var symbolExpr = regexp.MustCompile("[%$@/#=+\\-*&]")
var numExpr = regexp.MustCompile("\\d+")
var gearExpr = regexp.MustCompile("[*]")

func GetPartTotal(line string, neighbors ...string) int {
	var total int
	var indices []int
	for _, s := range neighbors {
		i := getSymbolIndices(s)
		indices = append(indices, i...)
	}
	numbers := numExpr.FindAllStringIndex(line, -1)
	if numbers == nil {
		return 0
	}
	for _, loc := range numbers {
		if checkIntersection(loc[0], loc[1], indices) {
			num, err := strconv.Atoi(line[loc[0]:loc[1]])
			if err != nil {
				slog.Error("Bad number", slog.Any("error", err))
			}
			total += num
		}
	}
	return total
}

func GetGearTotal(line string, neighbors ...string) int {
	var total int
	gears := gearExpr.FindAllStringIndex(line, -1)
	for _, loc := range gears {
		nearNumbers := getNumbersNear(loc[0], neighbors...)
		if len(nearNumbers) == 2 {
			total += nearNumbers[0] * nearNumbers[1]
		}
	}
	return total
}

func getSymbolIndices(line string) []int {
	symbols := symbolExpr.FindAllStringIndex(line, -1)
	var indices []int
	for _, i := range symbols {
		indices = append(indices, i[0]-1, i[0], i[0]+1)
	}
	return indices
}

func getNumbersNear(i int, neighbors ...string) []int {
	var nearNumbers []int
	for _, s := range neighbors {
		numbers := numExpr.FindAllStringIndex(s, -1)
		if numbers == nil {
			continue
		}
		for _, loc := range numbers {
			if checkIntersection(loc[0], loc[1], []int{i - 1, i, i + 1}) {
				num, err := strconv.Atoi(s[loc[0]:loc[1]])
				if err != nil {
					slog.Error("Bad number", slog.Any("error", err))
				}
				nearNumbers = append(nearNumbers, num)
			}
		}
	}
	return nearNumbers
}

func checkIntersection(start, end int, indices []int) bool {
	for i := start; i < end; i++ {
		if slices.Contains(indices, i) {
			return true
		}
	}
	return false
}
