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
	var prevLine, curLine, nextLine string
	for fileScanner.Scan() {
		prevLine = curLine
		curLine = nextLine
		nextLine = fileScanner.Text()
		total += engine.GetGearTotal(curLine, prevLine, curLine, nextLine)
	}
	total += engine.GetGearTotal(nextLine, curLine, nextLine)
	slog.Info("Answer", slog.Int("total", total))
}
