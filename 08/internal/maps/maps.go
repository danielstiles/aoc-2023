package maps

import (
	"regexp"
)

var nodeExpr = regexp.MustCompile("^(\\w+) = \\((\\w+), (\\w+)\\)$")
var left = byte('L')
var right = byte('R')

func GetNode(line string) *Node {
	groups := nodeExpr.FindStringSubmatch(line)
	if len(groups) != 4 {
		return nil
	}
	return &Node{
		Start: groups[1],
		Left:  groups[2],
		Right: groups[3],
	}
}

type Node struct {
	Start string
	Left  string
	Right string
}

type Map map[string]*Node

func NewMap() Map {
	return Map(make(map[string]*Node))
}

func (m Map) AddNode(n *Node) {
	m[n.Start] = n
}

func (m Map) Next(loc string, dir byte) string {
	if dir == left {
		return m[loc].Left
	}
	return m[loc].Right
}
