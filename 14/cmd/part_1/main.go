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
	g.Tilt(tilt.East)
	g.Print()
	slog.Info("Answer", slog.Int("weight", g.CalcWeight(tilt.North)))
}
