package main

import (
	"bufio"
	"os"

	"golang.org/x/exp/slog"

	"github.com/danielstiles/aoc-2023/08/internal/maps"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		slog.Error("Could not read file", slog.Any("error", err))
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	desertMap := maps.NewMap()
	fileScanner.Scan()
	directions := []byte(fileScanner.Text())
	for fileScanner.Scan() {
		node := maps.GetNode(fileScanner.Text())
		if node == nil {
			continue
		}
		desertMap.AddNode(node)
	}
	loc := "AAA"
	end := "ZZZ"
	var steps int
	for loc != end {
		loc = desertMap.Next(loc, directions[steps%(len(directions))])
		steps += 1
	}
	slog.Info("Answer", slog.Int("steps", steps))
}
