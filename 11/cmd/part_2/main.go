package main

import (
	"bufio"
	"log/slog"
	"os"

	"github.com/danielstiles/aoc-2023/11/internal/stars"
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
	chart := &stars.Chart{
		EmptyCost:    1000000,
		NonEmptyCost: 1,
	}
	for _, line := range lines {
		chart.AddLine([]byte(line))
	}
	slog.Info("Answer", slog.Int("total", chart.CalcDistances()))
}
