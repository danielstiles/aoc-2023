package main

import (
	"bufio"
	"os"

	"golang.org/x/exp/slog"

	"github.com/danielstiles/aoc-2023/06/internal/race"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		slog.Error("Could not read file", slog.Any("error", err))
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	product := 0
	fileScanner.Scan()
	times := race.GetNums(fileScanner.Text())
	fileScanner.Scan()
	distances := race.GetNums(fileScanner.Text())
	for i := range times {
		product *= race.GetRange(times[i], distances[i])
	}
	slog.Info("Answer", slog.Int("product", product))
}
