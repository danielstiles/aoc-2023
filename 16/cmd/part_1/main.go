package main

import (
	"bufio"
	"log/slog"
	"os"

	"github.com/danielstiles/aoc-2023/16/internal/mirrors"
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
	panel := &mirrors.Panel{}
	for _, line := range lines {
		panel.AddLine([]byte(line))
	}
	panel.SendCharge(0, 0, mirrors.East)
	slog.Info("Answer", slog.Int("total", panel.GetCharge()))
}
