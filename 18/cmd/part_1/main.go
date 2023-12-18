package main

import (
	"bufio"
	"log/slog"
	"os"

	"github.com/danielstiles/aoc-2023/18/internal/dig"
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
	var commands []dig.Command
	for _, line := range lines {
		commands = append(commands, dig.ParseLine([]byte(line), false))
	}
	digger := dig.NewDigger()
	for _, c := range commands {
		digger.Dig(c)
	}
	slog.Info("Answer", slog.Int("area", digger.Site.Area()))
}
