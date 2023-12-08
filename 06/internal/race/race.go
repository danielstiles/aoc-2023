package race

import (
	"math"
	"regexp"
	"strconv"
)

var numExpr = regexp.MustCompile("\\d+")

func GetNums(line string) []int {
	var nums []int
	for _, s := range numExpr.FindAllString(line, -1) {
		n, _ := strconv.Atoi(s)
		nums = append(nums, n)
	}
	return nums
}

func GetRange(time, distance int) int {
	interval := math.Sqrt(float64(time*time - 4*distance))
	base := int(math.Floor(interval))
	if time%2 == 0 {
		if base%2 == 0 {
			return base - 1
		}
	} else {
		if base%2 == 1 {
			return base - 1
		}
	}
	return base
}
