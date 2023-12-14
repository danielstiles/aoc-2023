package main

import (
	"bufio"
	"log/slog"
	"os"

	"github.com/danielstiles/aoc-2023/14/internal/tilt"
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
	g := &tilt.Grid{}
	for _, line := range lines {
		g.AddLine([]byte(line))
	}
	var prev []*tilt.Grid
	var i, pattern int
	iters := 1000000000
	for i = 0; i < iters && pattern == 0; i++ {
		if i%1000 == 0 {
			slog.Info("completed", slog.Int("loops", i))
		}
		curr := &tilt.Grid{}
		g.Copy(curr)
		prev = append(prev, curr)
		g.Tilt(tilt.North)
		g.Tilt(tilt.West)
		g.Tilt(tilt.South)
		g.Tilt(tilt.East)
		pattern = checkPrev(g, prev)
	}
	currCycle := i % pattern
	wantedCycle := iters % pattern
	for currCycle != wantedCycle {
		g.Tilt(tilt.North)
		g.Tilt(tilt.West)
		g.Tilt(tilt.South)
		g.Tilt(tilt.East)
		currCycle += 1
		currCycle %= pattern
	}
	slog.Info("Answer", slog.Int("weight", g.CalcWeight(tilt.North)))
}

func checkPrev(g *tilt.Grid, prev []*tilt.Grid) int {
	for i, p := range prev {
		if g.Equal(p) {
			return len(prev) - i
		}
	}
	return 0
}
