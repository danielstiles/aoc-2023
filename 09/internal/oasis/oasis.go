package oasis

import (
	"regexp"
	"strconv"
)

var numExpr = regexp.MustCompile("\\d+")

func ParseHistory(line string) []int {
	matches := numExpr.FindAllString(line)
	var nums []int
	for _, s := range matches {
		n, _ := strconv.Atoi(s)
		nums = append(nums, n)
	}
	return nums
}

func GetNext(pattern []int) int {
	var next int
	var done bool
	for !done {
		next += pattern[len(pattern)-1]
		var differences []int
		curr := pattern[0]
		done = true
		for _, n := range pattern[1:] {
			difference := n - curr
			if difference != 0 {
				done = false
			}
			differences = append(differences, difference)
			curr = n
		}
		pattern = differences
	}
	return next
}
