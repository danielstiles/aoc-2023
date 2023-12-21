package main

import (
	"bufio"
	"log/slog"
	"os"

	"github.com/danielstiles/aoc-2023/20/internal/signal"
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
	m := &signal.Machine{}
	for _, line := range lines {
		m.AddNode(signal.ParseLine([]byte(line)))
	}
	var high, low int
	for i := 0; i < 1000; i++ {
		newHigh, newLow := m.Run(false)
		high += newHigh
		low += newLow
	}
	slog.Info("Answer", slog.Int("total", high*low))
}
