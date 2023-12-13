package springs

import (
	"log/slog"
	"regexp"
	"strconv"
)

const (
	brokenSpring  = byte('#')
	unknownSpring = byte('?')
	workingSpring = byte('.')
)

const (
	brokenFlag  = 1
	unknownFlag = 2
)

var lineExpr = regexp.MustCompile("([.?#]+) ([0-9,]+)")
var numExpr = regexp.MustCompile("\\d+")

var cache = make(map[string]map[string]int)

type SpringRow struct {
	Broken      []int
	Arrangement []int
}

func ParseRow(line []byte, expand bool) SpringRow {
	matches := lineExpr.FindSubmatch(line)
	springs := matches[1]
	if expand {
		newSprings := make([]byte, len(springs)*5+4)
		for i := 0; i < 5; i++ {
			for j, spring := range springs {
				newSprings[i*len(springs)+i+j] = spring
			}
			if i != 4 {
				newSprings[(i+1)*len(springs)+i] = '?'
			}
		}
		springs = newSprings
	}
	arrangement := matches[2]
	if expand {
		newArrangement := make([]byte, len(arrangement)*5+4)
		for i := 0; i < 5; i++ {
			for j, arr := range arrangement {
				newArrangement[i*len(arrangement)+i+j] = arr
			}
			if i != 4 {
				newArrangement[(i+1)*len(arrangement)+i] = ','
			}
		}
		arrangement = newArrangement
	}
	ret := SpringRow{}
	for _, b := range springs {
		if b == brokenSpring {
			ret.Broken = append(ret.Broken, brokenFlag)
		} else if b == unknownSpring {
			ret.Broken = append(ret.Broken, unknownFlag)
		} else {
			ret.Broken = append(ret.Broken, 0)
		}
	}
	nums := numExpr.FindAll(arrangement, -1)
	for _, s := range nums {
		n, _ := strconv.Atoi(string(s))
		ret.Arrangement = append(ret.Arrangement, n)
	}
	return ret
}

func (s SpringRow) getMiddleUnknown(box []int) int {
	mid := (box[1] + box[0] - 1) / 2
	for i := 0; i < box[1]-mid; i++ {
		if (s.Broken[mid+i] & unknownFlag) != 0 {
			return mid + i
		}
		if mid-i >= box[0] && (s.Broken[mid-i]&unknownFlag) != 0 {
			return mid - i
		}
	}
	return -1
}

func (s SpringRow) Check() int {
	start := s.Print()
	startArr := s.PrintArrangement()
	if cache[start] == nil {
		cache[start] = make(map[string]int)
	}
	startCache := cache[start]
	if total, ok := startCache[startArr]; ok {
		return total
	}
	boxes := s.calcBoxes()
	boxes, placed := s.fill(boxes)
	if boxes == nil {
		if placed == -1 {
			startCache[startArr] = 0
			return 0
		}
		startCache[startArr] = 1
		return 1
	}
	if s.isPerfect(boxes) {
		startCache[startArr] = 1
		return 1
	}
	arr := s.Arrangement[placed:]
	if len(boxes) == 1 {
		if s.allUnknown(boxes[0]) {
			total := -1
			for _, a := range arr {
				total += a + 1
			}
			empty := boxes[0][1] - boxes[0][0] - total
			if empty < 0 {
				startCache[startArr] = 0
				return 0
			}
			bars := len(arr)
			startCache[startArr] = choose(empty+bars, empty)
			return startCache[startArr]
		}
		nextUnknown := s.getMiddleUnknown(boxes[0]) - boxes[0][0]
		if nextUnknown < 0 {
			slog.Info("shouldn't happen")
			startCache[startArr] = 0
			return 0
		}
		newBroken := make([]int, len(s.Broken[boxes[0][0]:boxes[0][1]]))
		newRow := SpringRow{
			Broken:      newBroken,
			Arrangement: arr,
		}
		copy(newBroken, s.Broken[boxes[0][0]:boxes[0][1]])
		newBroken[nextUnknown] = 0
		withWorking := newRow.Check()
		copy(newBroken, s.Broken[boxes[0][0]:boxes[0][1]])
		newBroken[nextUnknown] = brokenFlag
		withBroken := newRow.Check()
		startCache[startArr] = withWorking + withBroken
		return startCache[startArr]
	}
	var total int
	boxZero := fillBox(boxes[0][1]-boxes[0][0], arr)
	for i := 0; i <= boxZero; i++ {
		newZero := make([]int, boxes[0][1]-boxes[0][0])
		copy(newZero, s.Broken[boxes[0][0]:boxes[0][1]])
		rowZero := SpringRow{
			Broken:      newZero,
			Arrangement: arr[:i],
		}
		zeroCheck := rowZero.Check()
		rest := make([]int, len(s.Broken)-boxes[1][0])
		copy(rest, s.Broken[boxes[1][0]:])
		rowRest := SpringRow{
			Broken:      rest,
			Arrangement: arr[i:],
		}
		restCheck := rowRest.Check()
		total += zeroCheck * restCheck
	}
	startCache[startArr] = total
	return total
}

func (s SpringRow) calcBoxes() [][]int {
	var curLen int
	var boxes [][]int
	for i, spring := range s.Broken {
		if spring != 0 {
			curLen += 1
			continue
		}
		if curLen == 0 {
			continue
		}
		boxes = append(boxes, []int{i - curLen, i})
		curLen = 0
	}
	if curLen != 0 {
		boxes = append(boxes, []int{len(s.Broken) - curLen, len(s.Broken)})
	}
	return boxes
}

func (s SpringRow) fill(boxes [][]int) ([][]int, int) {
	if len(boxes) == 0 {
		if len(s.Arrangement) == 0 {
			return nil, 0
		}
		return nil, -1
	}
	var boxIndex, springIndex, boxEnd, placed int
	var unknown bool
	springIndex = boxes[boxIndex][0]
	boxEnd = boxes[boxIndex][1]
	for i, arr := range s.Arrangement {
		if boxIndex >= len(boxes) {
			return nil, -1
		}
		for boxEnd-springIndex < arr {
			for j := springIndex; j < boxEnd; j++ {
				if (s.Broken[j] & brokenFlag) != 0 {
					return nil, -1
				}
				s.Broken[j] = 0
			}
			boxIndex += 1
			if boxIndex >= len(boxes) {
				return nil, -1
			}
			springIndex = boxes[boxIndex][0]
			boxEnd = boxes[boxIndex][1]
		}
		if (s.Broken[springIndex] & brokenFlag) == 0 {
			unknown = true
			break
		}
		placed += 1
		for j := springIndex; j < springIndex+arr; j++ {
			if s.Broken[j] == 0 {
				return nil, -1
			}
			s.Broken[j] = brokenFlag
		}
		springIndex += arr
		if springIndex < boxEnd {
			if (s.Broken[springIndex] & brokenFlag) != 0 {
				return nil, -1
			}
			s.Broken[springIndex] = 0
			springIndex += 1
		}
		if springIndex >= boxEnd {
			boxIndex += 1
			if boxIndex >= len(boxes) {
				if i < len(s.Arrangement)-1 {
					return nil, -1
				}
				return nil, placed
			}
			springIndex = boxes[boxIndex][0]
			boxEnd = boxes[boxIndex][1]
		}
	}
	if !unknown {
		for j := springIndex; j < len(s.Broken); j++ {
			if (s.Broken[j] & brokenFlag) != 0 {
				return nil, -1
			}
			s.Broken[j] = 0
		}
		return nil, placed
	}

	newBoxes := [][]int{{springIndex, boxEnd}}
	for _, box := range boxes[boxIndex+1:] {
		newBoxes = append(newBoxes, box)
	}
	return newBoxes, placed
}

func (s SpringRow) isPerfect(boxes [][]int) bool {
	if len(boxes) == 0 {
		if len(s.Arrangement) == 0 {
			return true
		}
		return false
	}
	var boxIndex int
	springIndex := boxes[boxIndex][0]
	end := boxes[boxIndex][1]
	for i, arr := range s.Arrangement {
		if end-springIndex < arr {
			return false
		}
		for j := springIndex; j < springIndex+arr; j++ {
			if (s.Broken[j] & brokenFlag) == 0 {
				return false
			}
		}
		springIndex += arr
		if springIndex < end && (s.Broken[springIndex]&brokenFlag) != 0 {
			return false
		}
		springIndex += 1
		if springIndex >= end {
			boxIndex += 1
			if boxIndex == len(boxes) {
				if i == len(s.Arrangement)-1 {
					return true
				}
				return false
			}
			springIndex = boxes[boxIndex][0]
			end = boxes[boxIndex][1]
		}
	}
	return false
}

func (s SpringRow) allUnknown(box []int) bool {
	for _, spring := range s.Broken[box[0]:box[1]] {
		if (spring & unknownFlag) == 0 {
			return false
		}
	}
	return true
}

func fillBox(size int, arrangement []int) int {
	for i, arr := range arrangement {
		if size < arr {
			return i
		}
		size -= arr + 1
	}
	return len(arrangement)
}

func choose(n, k int) int {
	if k > n/2 {
		k = n - k
	}
	b := 1
	for i := 1; i <= k; i++ {
		b = (n - k + i) * b / i
	}
	return b
}

func (s SpringRow) Print() string {
	var ret string
	for _, spring := range s.Broken {
		if (spring & brokenFlag) != 0 {
			ret += "#"
		} else if (spring & unknownFlag) != 0 {
			ret += "?"
		} else {
			ret += "."
		}
	}
	return ret
}

func (s SpringRow) PrintArrangement() string {
	var ret string
	for _, arr := range s.Arrangement {
		if ret != "" {
			ret += ","
		}
		ret += strconv.Itoa(arr)
	}
	return ret
}
