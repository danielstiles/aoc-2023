package main

import (
	"bufio"
	"log/slog"
	"os"

	"github.com/danielstiles/aoc-2023/13/internal/mirrors"
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
	var grid [][]int
	var total int
	for _, line := range lines {
		if line == "" {
			total += mirrors.FindReflection(grid, false)
			grid = nil
			continue
		}
		grid = append(grid, mirrors.ParseLine([]byte(line)))
	}
	total += mirrors.FindReflection(grid, false)
	slog.Info("Answer", slog.Int("total", total))
}
