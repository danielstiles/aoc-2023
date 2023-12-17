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
	best := 0
	panel := &mirrors.Panel{}
	for _, line := range lines {
		panel.AddLine([]byte(line))
	}
	for row := 0; row < len(panel.Grid); row++ {
		panel.SendCharge(row, 0, mirrors.East)
		charge := panel.GetCharge()
		if charge > best {
			best = charge
		}
		panel = &mirrors.Panel{}
		for _, line := range lines {
			panel.AddLine([]byte(line))
		}
		panel.SendCharge(row, len(panel.Grid)-1, mirrors.West)
		charge = panel.GetCharge()
		if charge > best {
			best = charge
		}
		panel = &mirrors.Panel{}
		for _, line := range lines {
			panel.AddLine([]byte(line))
		}
	}
	if len(panel.Grid) != 0 {
		for col := 0; col < len(panel.Grid[0]); col++ {
			panel.SendCharge(0, col, mirrors.South)
			charge := panel.GetCharge()
			if charge > best {
				best = charge
			}
			panel = &mirrors.Panel{}
			for _, line := range lines {
				panel.AddLine([]byte(line))
			}
			panel.SendCharge(len(panel.Grid[0])-1, col, mirrors.North)
			charge = panel.GetCharge()
			if charge > best {
				best = charge
			}
			panel = &mirrors.Panel{}
			for _, line := range lines {
				panel.AddLine([]byte(line))
			}
		}
	}
	slog.Info("Answer", slog.Int("total", best))
}
