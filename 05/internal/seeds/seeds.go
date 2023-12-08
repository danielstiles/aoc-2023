package seeds

import (
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

func Convert(start []int, transform [][]int) []int {
	var end []int
	for _, seed := range start {
		var found bool
		for _, part := range transform {
			if seed >= part[1] && seed < (part[1]+part[2]) {
				end = append(end, seed-part[1]+part[0])
				found = true
				break
			}
		}
		if !found {
			end = append(end, seed)
		}
	}
	return end
}

func ConvertPairs(start [][]int, transform [][]int) [][]int {
	var end [][]int
	for _, pair := range start {
		ranges := [][]int{{pair[0], pair[1]}}
		found := true
		for found && len(ranges) > 0 {
			least := ranges[0][0]
			length := ranges[0][1]
			found = false
			for _, part := range transform {
				if least < part[1]+part[2] && least+length > part[1] {
					if least < part[1] {
						ranges = append(ranges, []int{least, part[1] - least})
						least = part[1]
						length = least + length - part[1]
					}
					if least+length > part[1]+part[2] {
						ranges = append(ranges, []int{part[1] + part[2], least + length - (part[1] + part[2])})
						length = part[1] + part[2] - least
					}
					end = append(end, []int{least - part[1] + part[0], length})
					found = true
					ranges = ranges[1:]
					break
				}
			}
		}
		for _, part := range ranges {
			end = append(end, part)
		}
	}
	return end
}
