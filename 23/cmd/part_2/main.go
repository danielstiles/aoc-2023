package main

import (
	"bufio"
	"log/slog"
	"os"

	"github.com/danielstiles/aoc-2023/23/internal/hikes"
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
	maze := &hikes.Maze{}
	for _, line := range lines {
		maze.AddLine([]byte(line))
	}
	var start, end hikes.Pos
	start.Row = 0
	start.Steps = 1
	end.Row = len(lines) - 1
	end.Steps = 1
	for col, t := range []byte(lines[start.Row]) {
		if t == hikes.Empty {
			start.Col = col
			break
		}
	}
	for col, t := range []byte(lines[end.Row]) {
		if t == hikes.Empty {
			end.Col = col
			break
		}
	}
	best := maze.FindLongestPath(start, end, false)
	slog.Info("Answer", slog.Int("length", best-1))
}
