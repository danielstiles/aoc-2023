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
	g := &garden.Garden{}
	for _, line := range lines {
		g.AddLine([]byte(line))
	}
	g.Tiles[65][65] = byte('S')
	possibilities := 0
	for i := 0; i < 64; i++ {
		possibilities = g.Iterate()
	}
	slog.Info("Answer", slog.Int("total", possibilities))
}
