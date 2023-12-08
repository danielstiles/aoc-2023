package main

import (
	"bufio"
	"os"

	"golang.org/x/exp/slices"
	"golang.org/x/exp/slog"

	"github.com/danielstiles/aoc-2023/05/internal/seeds"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		slog.Error("Could not read file", slog.Any("error", err))
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var curr []int
	var block [][]int
	var inBlock bool
	for fileScanner.Scan() {
		line := fileScanner.Text()
		nums := seeds.GetNums(line)
		if curr == nil {
			curr = nums
			continue
		}
		if nums != nil {
			inBlock = true
			block = append(block, nums)
			continue
		}
		if !inBlock {
			continue
		}
		curr = seeds.Convert(curr, block)
		block = nil
		inBlock = false
	}
	if inBlock {
		curr = seeds.Convert(curr, block)
		block = nil
		inBlock = false
	}
	slices.Sort(curr)
	slog.Info("Answer", slog.Int("least", curr[0]))
}
