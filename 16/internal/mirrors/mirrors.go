package mirrors

const (
	ground    = byte('.')
	hSplitter = byte('-')
	vSplitter = byte('|')
	fMirror   = byte('/')
	bMirror   = byte('\\')
)

const (
	North = 1
	East  = 2
	South = 4
	West  = 8
)

type Panel struct {
	Grid    [][]byte
	Powered [][]int
}

func (p *Panel) AddLine(line []byte) {
	p.Grid = append(p.Grid, line)
	p.Powered = append(p.Powered, make([]int, len(line)))
	return
}

func (p *Panel) SendCharge(row, col, dir int) {
	if len(p.Grid) == 0 {
		return
	}
	for row >= 0 && row < len(p.Grid) && col >= 0 && col < len(p.Grid[0]) {
		if (p.Powered[row][col] & dir) != 0 {
			return
		}
		p.Powered[row][col] |= dir
		switch p.Grid[row][col] {
		case ground, fMirror, bMirror:
			row, col, dir = move(row, col, p.Grid[row][col], dir)
		case hSplitter:
			switch dir {
			case North, South:
				dupeRow, dupeCol, dupeDir := move(row, col, ground, East)
				p.SendCharge(dupeRow, dupeCol, dupeDir)
				row, col, dir = move(row, col, ground, West)
			case East, West:
				row, col, dir = move(row, col, ground, dir)
			}
		case vSplitter:
			switch dir {
			case East, West:
				dupeRow, dupeCol, dupeDir := move(row, col, ground, North)
				p.SendCharge(dupeRow, dupeCol, dupeDir)
				row, col, dir = move(row, col, ground, South)
			case North, South:
				row, col, dir = move(row, col, ground, dir)
			}
		}
	}
}

func move(row, col int, tile byte, dir int) (int, int, int) {
	switch tile {
	case ground:
		switch dir {
		case North:
			row -= 1
		case East:
			col += 1
		case South:
			row += 1
		case West:
			col -= 1
		}
	case fMirror:
		switch dir {
		case North:
			col += 1
			dir = East
		case East:
			row -= 1
			dir = North
		case South:
			col -= 1
			dir = West
		case West:
			row += 1
			dir = South
		}
	case bMirror:
		switch dir {
		case North:
			col -= 1
			dir = West
		case East:
			row += 1
			dir = South
		case South:
			col += 1
			dir = East
		case West:
			row -= 1
			dir = North
		}
	}
	return row, col, dir
}

func (p *Panel) GetCharge() int {
	var total int
	for _, arr := range p.Powered {
		for _, power := range arr {
			if power != 0 {
				total += 1
			}
		}
	}
	return total
}
