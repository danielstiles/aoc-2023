package blocks

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

type Vector3 struct {
	X int
	Y int
	Z int
}

type Block struct {
	Start         Vector3
	End           Vector3
	allSupporting []*Block
	allFalling    []*Block
	supporting    []*Block
	supportedBy   []*Block
}

type Space struct {
	Occupied [][][]bool
	Blocks   []*Block
}

var numExpr = regexp.MustCompile("\\d+")

func ParseBlock(line []byte) *Block {
	nums := numExpr.FindAll(line, -1)
	if len(nums) != 6 {
		return nil
	}
	x1, _ := strconv.Atoi(string(nums[0]))
	y1, _ := strconv.Atoi(string(nums[1]))
	z1, _ := strconv.Atoi(string(nums[2]))
	x2, _ := strconv.Atoi(string(nums[3]))
	y2, _ := strconv.Atoi(string(nums[4]))
	z2, _ := strconv.Atoi(string(nums[5]))
	return &Block{
		Start: Vector3{
			X: min(x1, x2),
			Y: min(y1, y2),
			Z: min(z1, z2),
		},
		End: Vector3{
			X: max(x1, x2),
			Y: max(y1, y2),
			Z: max(z1, z2),
		},
	}
}

func (b *Block) Intersects(other *Block) bool {
	return b.Start.X <= other.End.X && other.Start.X <= b.End.X &&
		b.Start.Y <= other.End.Y && other.Start.Y <= b.End.Y &&
		b.Start.Z <= other.End.Z && other.Start.Z <= b.End.Z
}

func (b *Block) NumFalling() int {
	return len(b.GetAllFalling())
}

func (b *Block) GetAllFalling() []*Block {
	if b.allFalling != nil {
		return b.allFalling
	}
	var falling []*Block
	for _, block := range b.GetAllSupporting() {
		falling = append(falling, block)
	}
	changed := true
	for changed {
		changed = false
		for i := 0; i < len(falling); i++ {
			del := false
			for _, block2 := range falling[i].supportedBy {
				if block2 != b && !slices.Contains(falling, block2) {
					del = true
					break
				}
			}
			if del {
				changed = true
				falling = slices.Delete(falling, i, i+1)
				i -= 1
			}
		}
	}
	b.allFalling = falling
	return b.allFalling
}

func (b *Block) GetAllSupporting() []*Block {
	if b.allSupporting != nil {
		return b.allSupporting
	}
	for _, block := range b.supporting {
		if !slices.Contains(b.allSupporting, block) {
			b.allSupporting = append(b.allSupporting, block)
		}
		for _, block2 := range block.GetAllSupporting() {
			if !slices.Contains(b.allSupporting, block2) {
				b.allSupporting = append(b.allSupporting, block2)
			}
		}
	}
	return b.allSupporting
}

func (b *Block) Print() string {
	return fmt.Sprintf("%d,%d,%d~%d,%d,%d",
		b.Start.X, b.Start.Y, b.Start.Z,
		b.End.X, b.End.Y, b.End.Z)
}

func MaxPos(blocks []*Block) Vector3 {
	maxPos := Vector3{}
	for _, block := range blocks {
		maxPos.X = max(maxPos.X, block.Start.X, block.End.X)
		maxPos.Y = max(maxPos.Y, block.Start.Y, block.End.Y)
		maxPos.Z = max(maxPos.Z, block.Start.Z, block.End.Z)
	}
	return maxPos
}

func NewSpace(dims Vector3) *Space {
	var space [][][]bool
	for i := 0; i <= dims.X; i++ {
		var newRow [][]bool
		for j := 0; j <= dims.Y; j++ {
			newCol := []bool{true}
			for k := 1; k <= dims.Z; k++ {
				newCol = append(newCol, false)
			}
			newRow = append(newRow, newCol)
		}
		space = append(space, newRow)
	}
	return &Space{
		Occupied: space,
	}
}

func (s *Space) AddBlocks(blocks []*Block) {
	slices.SortFunc(blocks, func(a, b *Block) int { return min(a.Start.Z, a.End.Z) - min(b.Start.Z, b.End.Z) })
	for _, block := range blocks {
		s.PlaceBlock(block)
	}
}

func (s *Space) PlaceBlock(block *Block) {
	for !s.CheckOccupied(block) {
		block.Start.Z -= 1
		block.End.Z -= 1
	}
	for _, block2 := range s.Blocks {
		if block2.Intersects(block) {
			block2.supporting = append(block2.supporting, block)
			block.supportedBy = append(block.supportedBy, block2)
		}
	}
	block.Start.Z += 1
	block.End.Z += 1
	for i := block.Start.X; i <= block.End.X; i++ {
		for j := block.Start.Y; j <= block.End.Y; j++ {
			for k := block.Start.Z; k <= block.End.Z; k++ {
				s.Occupied[i][j][k] = true
			}
		}
	}
	s.Blocks = append(s.Blocks, block)
}

func (s *Space) CheckOccupied(block *Block) bool {
	for i := block.Start.X; i <= block.End.X; i++ {
		for j := block.Start.Y; j <= block.End.Y; j++ {
			for k := block.Start.Z; k <= block.End.Z; k++ {
				if s.Occupied[i][j][k] {
					return true
				}
			}
		}
	}
	return false
}
