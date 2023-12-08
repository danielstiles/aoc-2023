package main

import (
	"bufio"
	"os"

	"golang.org/x/exp/slog"

	"github.com/danielstiles/aoc-2023/04/internal/lotto"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		slog.Error("Could not read file", slog.Any("error", err))
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	cardCount := make(map[int]int)
	var total, index int
	for fileScanner.Scan() {
		index += 1
		cardCount[index] += 1
		winners := lotto.GetWinners(fileScanner.Text())
		for i := index + 1; i <= index+winners; i++ {
			cardCount[i] += cardCount[index]
		}
		total += cardCount[index]
	}
	slog.Info("Answer", slog.Int("total", total))
}
