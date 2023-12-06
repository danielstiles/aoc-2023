package main

import (
	"bufio"
	"os"

	"golang.org/x/exp/slog"

	"github.com/danielstiles/aoc-2023/02/internal/bag"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		slog.Error("Could not read file", slog.Any("error", err))
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	total := 0
	for fileScanner.Scan() {
		total += bag.CheckGame(fileScanner.Text())
	}
	slog.Info("Answer", slog.Int("total", total))
}
