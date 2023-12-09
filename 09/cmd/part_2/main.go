package main

import (
	"bufio"
	"os"

	"golang.org/x/exp/slices"
	"golang.org/x/exp/slog"

	"github.com/danielstiles/aoc-2023/09/internal/oasis"
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
		line := fileScanner.Text()
		history := oasis.ParseHistory(line)
		slices.Reverse(history)
		total += oasis.GetNext(history)
	}
	slog.Info("Answer", slog.Int("total", total))
}
