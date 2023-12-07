package engine

import "golang.org/x/exp/slices"

var symbols = []byte{'%', '$', '@', '/', '#', '=', '+', '-', '*', '&'}

func GetIndices(line string) map[int]struct{} {
	var indices = make(map[int]struct{})
	for i, c := range []byte(line) {
		if slices.Contains(symbols, c) {
			if i > 0 {
				indices[i-1] = struct{}{}
			}
			indices[i] = struct{}{}
			if i < len(line)-1 {
				indices[i+1] = struct{}{}
			}
		}
	}
	return indices
}

func GetPartTotal(line string, indexSets ...map[int]struct{}) int {
	var total int
	var indices = make(map[int]struct{})
	mergeIndices(indices, indexSets...)
	for i, _ := range indices {
		total += int(line[i]) - 48
	}
	return total
}

func mergeIndices(into map[int]struct{}, from ...map[int]struct{}) {
	for _, m := range from {
		for i, _ := range m {
			into[i] = struct{}{}
		}
	}
}
