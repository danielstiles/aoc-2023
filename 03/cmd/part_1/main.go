package main

import (
	"bufio"
	"os"

	"golang.org/x/exp/slog"

	"github.com/danielstiles/aoc-2023/03/internal/engine"
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
	var prevIndices, curIndices, nextIndices map[int]struct{}
	var curLine, nextLine string
	for fileScanner.Scan() {
		prevIndices = curIndices
		curLine = nextLine
		curIndices = nextIndices
		nextLine = fileScanner.Text()
		nextIndices = engine.GetIndices(nextLine)
		total += engine.GetPartTotal(curLine, prevIndices, curIndices, nextIndices)
	}
	total += engine.GetPartTotal(nextLine, curIndices, nextIndices)
	slog.Info("Answer", slog.Int("total", total))
}
