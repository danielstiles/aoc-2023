package calibration

import (
	"regexp"
	"unicode/utf8"

	"golang.org/x/exp/slog"
)

func GetCalibration(line string) int {
	out := 0
	var first, last bool
	l := len(line)
	for i, b := range line {
		b2 := line[l-i-1]
		if !first && b >= '0' && b <= '9' {
			out += (int(b) - 48) * 10
			first = true
		}
		if !last && b2 >= '0' && b2 <= '9' {
			out += int(b2) - 48
			last = true
		}
		if first && last {
			break
		}
	}
	return out
}

var digits = map[string]int{
	"0":     0,
	"zero":  0,
	"1":     1,
	"one":   1,
	"2":     2,
	"two":   2,
	"3":     3,
	"three": 3,
	"4":     4,
	"four":  4,
	"5":     5,
	"five":  5,
	"6":     6,
	"six":   6,
	"7":     7,
	"seven": 7,
	"8":     8,
	"eight": 8,
	"9":     9,
	"nine":  9,
}

var forward, backward *regexp.Regexp

func GetV2Calibration(line string) int {
	if forward == nil {
		buildRegex()
	}
	first := forward.FindString(line)
	if first == "" {
		return 0
	}
	last := backward.FindString(reverseString(line))
	return digits[first]*10 + digits[reverseString(last)]
}

func buildRegex() {
	r := ""
	first := true
	for name := range digits {
		if !first {
			r += "|"
		}
		first = false
		r += name
	}
	var err error
	forward, err = regexp.Compile(r)
	if err != nil {
		slog.Error("Could not compile forward regex", slog.Any("error", err))
	}
	backward, err = regexp.Compile(reverseString(r))
	if err != nil {
		slog.Error("Could not compile backward regex", slog.Any("error", err))
	}
}

func reverseString(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}
