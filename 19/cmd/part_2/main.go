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
	for _, line := range lines {
		if line == "" {
			break
		}
		p.AddRule(machine.ParseRule([]byte(line)))
	}
	ranges := p.GetAccepatableRanges()
	total := 0
	for _, rule := range ranges {
		var minX, maxX, minM, maxM, minA, maxA, minS, maxS = 1, 4000, 1, 4000, 1, 4000, 1, 4000
		var fail bool
		for _, cond := range rule.Conditions {
			var min, max *int
			switch cond.Variable {
			case machine.X:
				min = &minX
				max = &maxX
			case machine.M:
				min = &minM
				max = &maxM
			case machine.A:
				min = &minA
				max = &maxA
			case machine.S:
				min = &minS
				max = &maxS
			}
			switch cond.Condition {
			case machine.Greater:
				if cond.Level+1 > *min {
					*min = cond.Level + 1
				}
			case machine.Less:
				if cond.Level-1 < *max {
					*max = cond.Level - 1
				}
			case machine.GreaterEqual:
				if cond.Level > *min {
					*min = cond.Level
				}
			case machine.LessEqual:
				if cond.Level < *max {
					*max = cond.Level
				}
			}
			if *min > *max {
				fail = true
				break
			}
		}
		if !fail {
			total += (maxX - minX + 1) * (maxM - minM + 1) * (maxA - minA + 1) * (maxS - minS + 1)
		}
	}
	slog.Info("Answer", slog.Int("total", total))
}
