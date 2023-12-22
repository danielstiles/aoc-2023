package main

import (
	"bufio"
	"log/slog"
	"os"

	"github.com/danielstiles/aoc-2023/22/internal/blocks"
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
	var b []*blocks.Block
	for _, line := range lines {
		b = append(b, blocks.ParseBlock([]byte(line)))
	}
	dims := blocks.MaxPos(b)
	s := blocks.NewSpace(dims)
	s.AddBlocks(b)
	total := 0
	for _, block := range b {
		total += block.NumFalling()
	}
	slog.Info("Answer", slog.Int("total", total))
}
