package bag

import (
	"regexp"
	"strconv"

	"golang.org/x/exp/slog"
)

var maxPick = map[string]int{
	"Red":   12,
	"Green": 13,
	"Blue":  14,
}

var gameExpr, drawExpr, pickExpr *regexp.Regexp

func CheckGame(line string) int {
	if gameExpr == nil {
		buildRegex()
	}
	parts := gameExpr.FindStringSubmatch(line)
	game, err := strconv.Atoi(parts[1])
	if err != nil {
		slog.Error("Bad game number", slog.Any("error", err))
	}
	for i, draw := range parts[2:] {
		if i > 0 {
			draw = draw[2:]
		}
		picks := drawExpr.FindStringSubmatch(draw)
		for j, pick := range picks[1:] {
			if j > 0 {
				pick = pick[2:]
			}
			if !checkPick(pick) {
				return 0
			}
		}
	}
	return game
}

func buildRegex() {
	gameExpr = regexp.MustCompile("^Game (\\d+): ([^;]+)(; [^;]+)*$")
	drawExpr = regexp.MustCompile("^(\\d+ \\w+)(, \\d+ \\w+)*$")
	pickExpr = regexp.MustCompile("^(\\d+) (Red|Green|Blue)$")
}

func checkPick(pick string) bool {
	parts := pickExpr.FindStringSubmatch(pick)
	num, err := strconv.Atoi(parts[1])
	if err != nil {
		slog.Error("Bad pick number", slog.Any("error", err))
	}
	return num <= maxPick[parts[2]]
}
