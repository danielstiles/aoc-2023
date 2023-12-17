package main

import (
	"bufio"
	"log/slog"
	"os"

	"github.com/danielstiles/aoc-2023/17/internal/factory"
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
	city := &factory.City{MaxLength: 3}
	for _, line := range lines {
		city.AddLine([]byte(line))
	}
	pathLen := city.FindPath(0, 0, len(city.Costs)-1, len(city.Costs[0])-1)
	slog.Info("Answer", slog.Int("length", pathLen))
}
