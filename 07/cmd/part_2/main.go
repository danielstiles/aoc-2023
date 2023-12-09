package main

import (
	"bufio"
	"os"

	"golang.org/x/exp/slices"
	"golang.org/x/exp/slog"

	"github.com/danielstiles/aoc-2023/07/internal/hands"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		slog.Error("Could not read file", slog.Any("error", err))
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var lines []hands.Line
	for fileScanner.Scan() {
		lines = append(lines, hands.ParseLine(fileScanner.Text()))
	}
	slices.SortFunc(lines, hands.GetCompareFunc(true))
	var total int
	for i, line := range lines {
		total += (i + 1) * line.Bid
	}
	slog.Info("Answer", slog.Int("total", total))
}
