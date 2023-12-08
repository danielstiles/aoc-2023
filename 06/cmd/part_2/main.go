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
	fileScanner.Scan()
	time := race.GetNum(fileScanner.Text())
	fileScanner.Scan()
	distance := race.GetNum(fileScanner.Text())
	product := race.GetRange(time, distance)
	slog.Info("Answer", slog.Int("product", product))
}
