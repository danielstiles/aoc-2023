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
	startChar := byte('A')
	endChar := byte('Z')
	var locs [][]byte
	for fileScanner.Scan() {
		node := maps.GetNode(fileScanner.Text())
		if node == nil {
			continue
		}
		loc := []byte(node.Start)
		if loc[len(loc)-1] == startChar {
			locs = append(locs, loc)
		}
		desertMap.AddNode(node)
	}
	steps := 1
	var done bool
	cycleCheck := make(map[int]map[string][]int)
	cycles := make(map[int][]int)
	var foundCycles int
	for !done {
		for i, loc := range locs {
			if len(loc) == 0 {
				continue
			}
			if cycleCheck[i] == nil {
				cycleCheck[i] = make(map[string][]int)
			}
			nextLoc := []byte(desertMap.Next(string(loc), directions[(steps-1)%(len(directions))]))
			if nextLoc[len(nextLoc)-1] == endChar {
				for _, prevSteps := range cycleCheck[i][string(nextLoc)] {
					if (steps-prevSteps)%(len(directions)) == 0 {
						cycleLen := (steps - prevSteps) / len(directions)
						cycles[i] = []int{cycleLen}
						for _, j := range cycleCheck[i][string(nextLoc)] {
							cycles[i] = append(cycles[i], j%cycleLen)
						}
						nextLoc = []byte("")
						foundCycles += 1
						break
					}
				}
				if len(nextLoc) != 0 {
					cycleCheck[i][string(nextLoc)] = append(cycleCheck[i][string(nextLoc)], steps)
				}
			}
			locs[i] = nextLoc
		}
		done = foundCycles == len(locs)
		steps += 1
	}
	var possible []int
	var nextPossible []int
	for _, cycle := range cycles {
		if possible == nil {
			possible = cycle
			continue
		}
		cycleLen := possible[0]
		var expand bool
		if possible[0]%cycle[0] != 0 {
			cycleLen = LCM(possible[0], cycle[0])
			expand = true
		}
		nextPossible = []int{cycleLen}
		gcd := GCD(possible[0], cycle[0])
		for _, i := range possible[1:] {
			for _, j := range cycle[1:] {
				if (i-j)%(gcd*len(directions)) == 0 {
					if !expand {
						nextPossible = append(nextPossible, i)
					}
					cycleI := i / len(directions)
					cycleJ := j / len(directions)
					r := cycleI % gcd
					if cycleJ%gcd == r {
						cycleI = cycleI / gcd
						cycleJ = cycleJ / gcd
						meeting := CombineRemainders(cycleI, cycleJ, possible[0]/gcd, cycle[0]/gcd)
						meeting = meeting*gcd + r
						meeting = meeting*len(directions) + (i % len(directions))
						nextPossible = append(nextPossible, meeting)
					}
				}
			}
		}
		possible = nextPossible
	}
	slog.Info("Answer", slog.Int("cycle", possible[0]*len(directions)))
	for _, i := range possible[1:] {
		slog.Info("Answer", slog.Int("steps", i))
	}
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func CombineRemainders(r1, r2, m1, m2 int) int {
	p1 := 1
	p2 := 1
	if r2 != 0 {
		for i := 0; i < m2-1; i++ {
			p1 = (p1 * m1) % (m1 * m2)
		}
	}
	if r1 != 0 {
		for i := 0; i < m1-1; i++ {
			p2 = (p2 * m2) % (m1 * m2)
		}
	}
	return (r1*p2 + r2*p1) % (m1 * m2)
}
