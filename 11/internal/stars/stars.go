package stars

const (
	empty = byte('.')
	star  = byte('#')
)

type Chart struct {
	EmptyCost    int
	NonEmptyCost int
	Stars        [][]int
	RowEmpty     []bool
	ColEmpty     []bool
}

func (c *Chart) AddLine(line []byte) {
	var first, empty bool
	empty = true
	if c.ColEmpty == nil {
		first = true
	}
	for i, b := range line {
		if first {
			c.ColEmpty = append(c.ColEmpty, true)
		}
		if b == star {
			empty = false
			c.ColEmpty[i] = false
			c.Stars = append(c.Stars, []int{len(c.RowEmpty), i})
		}
	}
	c.RowEmpty = append(c.RowEmpty, empty)
}

func (c *Chart) CalcDistances() int {
	var total int
	for i, a := range c.Stars {
		for _, b := range c.Stars[i+1:] {
			total += c.calcDistance(a, b)
		}
	}
	return total
}

func (c *Chart) calcDistance(a, b []int) int {
	var distance, min, max int
	if a[0] > b[0] {
		min = b[0]
		max = a[0]
	} else {
		min = a[0]
		max = b[0]
	}
	for i := min; i < max; i++ {
		if c.RowEmpty[i] {
			distance += c.EmptyCost
		} else {
			distance += c.NonEmptyCost
		}
	}
	if a[1] > b[1] {
		min = b[1]
		max = a[1]
	} else {
		min = a[1]
		max = b[1]
	}
	for i := min; i < max; i++ {
		if c.ColEmpty[i] {
			distance += c.EmptyCost
		} else {
			distance += c.NonEmptyCost
		}
	}
	return distance
}
