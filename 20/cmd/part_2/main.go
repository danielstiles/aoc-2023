package main

import (
	"bufio"
	"fmt"
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
	rx := signal.FlipFlop{}
	rx.Name = "rx"
	m.AddNode(&rx)
	var count int
	var found bool
	for count = 0; count < 20000; count++ {
		m.Run(false)
		if found != rx.State {
			slog.Info("output", slog.Int("iterations", count+1), slog.String("state", fmt.Sprintf("%049b", m.GetState())))
			if found {
				break
			}
			found = true
		}
	}
	slog.Info("Answer", slog.Int("total", count))
}
