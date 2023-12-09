package hands

import (
	"regexp"
	"strconv"
	"strings"
)

var order = "23456789TJQKA"
var orderJokers = "J23456789TQKA"
var joker = byte('J')
var lineExpr = regexp.MustCompile("^([AKQJT2-9]{5}) (\\d+)$")

const (
	HighCard = iota
	Pair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Line struct {
	Hand []byte
	Bid  int
}

func ParseLine(line string) Line {
	groups := lineExpr.FindStringSubmatch(line)
	hand := groups[1]
	bidStr := groups[2]
	bid, _ := strconv.Atoi(bidStr)
	return Line{
		Hand: []byte(hand),
		Bid:  bid,
	}
}

func GetCompareFunc(jokers bool) func(a, b Line) int {
	var o string
	if jokers {
		o = orderJokers
	} else {
		o = order
	}
	return func(a, b Line) int {
		typeA := GetType(a.Hand, jokers)
		typeB := GetType(b.Hand, jokers)
		if typeA != typeB {
			return typeA - typeB
		}
		for i := range a.Hand {
			if a.Hand[i] != b.Hand[i] {
				valueA := strings.Index(o, string(a.Hand[i]))
				valueB := strings.Index(o, string(b.Hand[i]))
				return valueA - valueB
			}
		}
		return 0
	}
}

func GetType(hand []byte, jokers bool) int {
	cards := make(map[byte]int)
	for _, c := range hand {
		cards[c] += 1
	}
	var most, second, wild int
	for card, count := range cards {
		if jokers && card == joker {
			wild = count
		} else if count > most {
			second = most
			most = count
		} else if count > second {
			second = count
		}
	}
	most += wild
	switch most {
	case 1:
		return HighCard
	case 2:
		if second == 2 {
			return TwoPair
		}
		return Pair
	case 3:
		if second == 2 {
			return FullHouse
		}
		return ThreeOfAKind
	case 4:
		return FourOfAKind
	case 5:
		return FiveOfAKind
	}
	return 0
}
