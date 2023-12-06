package bag

import (
	"regexp"
	"strconv"

	"golang.org/x/exp/slog"
)

var maxPick = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

var gameExpr, pickExpr *regexp.Regexp

func CheckGame(line string) int {
	if gameExpr == nil {
		buildRegex()
	}
	parts := gameExpr.FindStringSubmatch(line)
	game, err := strconv.Atoi(parts[1])
	if err != nil {
		slog.Error("Bad game number", slog.Any("error", err))
	}
	picks := pickExpr.FindAllStringSubmatch(line, -1)
	for _, pick := range picks {
		n, err := strconv.Atoi(pick[1])
		if err != nil {
			slog.Error("Bad pick number", slog.Any("error", err))
		}
		if !checkPick(n, pick[2]) {
			return 0
		}
	}
	return game
}

func GetPower(line string) int {
	if gameExpr == nil {
		buildRegex()
	}
	var required = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	picks := pickExpr.FindAllStringSubmatch(line, -1)
	for _, pick := range picks {
		n, err := strconv.Atoi(pick[1])
		if err != nil {
			slog.Error("Bad pick number", slog.Any("error", err))
		}
		if required[pick[2]] < n {
			required[pick[2]] = n
		}
	}
	return required["red"] * required["green"] * required["blue"]
}

func buildRegex() {
	gameExpr = regexp.MustCompile("^Game (\\d+)")
	pickExpr = regexp.MustCompile("(\\d+) (red|green|blue)")
}

func checkPick(num int, color string) bool {
	return num <= maxPick[color]
}
