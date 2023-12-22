package main

import (
	"bufio"
	"log/slog"
	"os"

	"github.com/danielstiles/aoc-2023/21/internal/garden"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		slog.Error("Could not read file", slog.Any("error", err))
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	var gardens []*garden.Garden
	for i := 0; i < 9; i++ {
		gardens = append(gardens, &garden.Garden{})
	}
	for _, line := range lines {
		for _, g := range gardens {
			bytes := make([]byte, len([]byte(line)))
			copy(bytes, []byte(line))
			g.AddLine(bytes)
		}
	}
	gardens[0].Tiles[65][65] = byte('S')
	gardens[1].Tiles[65][0] = byte('S')
	gardens[2].Tiles[65][130] = byte('S')
	gardens[3].Tiles[0][65] = byte('S')
	gardens[4].Tiles[130][65] = byte('S')
	gardens[5].Tiles[0][0] = byte('S')
	gardens[6].Tiles[0][130] = byte('S')
	gardens[7].Tiles[130][0] = byte('S')
	gardens[8].Tiles[130][130] = byte('S')
	for i := 1; i <= 64; i++ {
		gardens[0].Iterate()
	}
	for i := 1; i <= 129; i++ {
		for j := 1; j <= 4; j++ {
			gardens[j].Iterate()
		}
	}
	for i := 1; i <= 63; i++ {
		for j := 5; j <= 8; j++ {
			gardens[j].Iterate()
		}
	}
	round0 := make(map[int][]int)
	for i := 0; i < 394; i++ {
		for j, g := range gardens {
			count := g.Iterate()
			if i%131 == 0 {
				round0[j] = append(round0[j], count)
			}
		}
	}
	total := 0
	limit := 202300
	for i := 0; i <= limit; i++ {
		index := limit - i
		if index > 3 {
			index = index%2 + 2
		}
		switch i {
		case 0:
			total += round0[0][index]
		default:
			total += round0[1][index]
			total += round0[2][index]
			total += round0[3][index]
			total += round0[4][index]
			total += round0[5][index] * i
			total += round0[6][index] * i
			total += round0[7][index] * i
			total += round0[8][index] * i
		}
	}

	slog.Info("Answer", slog.Int("total", total))
}
