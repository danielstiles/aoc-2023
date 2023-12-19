package main

import (
	"bufio"
	"log/slog"
	"os"

	"github.com/danielstiles/aoc-2023/19/internal/machine"
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
	p := machine.NewProgram()
	var programDone bool
	var accepted []machine.Part
	for _, line := range lines {
		if line == "" {
			programDone = true
			continue
		}
		if !programDone {
			p.AddRule(machine.ParseRule([]byte(line)))
			continue
		}
		part := machine.ParsePart([]byte(line))
		accept := p.Process(part)
		if accept {
			accepted = append(accepted, part)
		}
	}
	var total int
	for _, part := range accepted {
		total += part.X + part.M + part.A + part.S
	}
	slog.Info("Answer", slog.Int("total", total))
}
