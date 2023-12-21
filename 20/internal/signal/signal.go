package signal

import (
	"regexp"
	"slices"
)

var lineExpr = regexp.MustCompile("([\\w&%]+) -> ([\\w ,]+)")
var targetExpr = regexp.MustCompile("\\w+")

const (
	flipFlop   = byte('%')
	conjuction = byte('&')
)

func ParseLine(line []byte) Node {
	if len(line) == 0 {
		return nil
	}
	groups := lineExpr.FindSubmatch(line)
	targets := targetExpr.FindAll(groups[2], -1)
	var node Node
	switch groups[1][0] {
	case flipFlop:
		node = &FlipFlop{baseNode: baseNode{Name: string(groups[1][1:])}}
	case conjuction:
		node = &Conjunction{baseNode: baseNode{Name: string(groups[1][1:])}}
	default:
		node = &Broadcast{baseNode: baseNode{Name: string(groups[1])}}
	}
	for _, target := range targets {
		node.AddTarget(string(target))
	}
	return node
}

const (
	button    = "button"
	startNode = "broadcaster"
)

type Machine struct {
	Nodes    map[string]Node
	complete bool
}

func (m *Machine) AddNode(node Node) {
	if m.Nodes == nil {
		m.Nodes = make(map[string]Node)
	}
	if node == nil {
		return
	}
	m.Nodes[node.GetName()] = node
}

func (m *Machine) AddInputs() {
	for _, node := range m.Nodes {
		for _, target := range node.GetTargets() {
			if targetNode, ok := m.Nodes[target]; ok {
				targetNode.AddInput(node.GetName())
			}
		}
	}
}

func (m *Machine) Run(high bool) (int, int) {
	if !m.complete {
		m.AddInputs()
		m.complete = true
	}
	highPulses := 0
	lowPulses := 0
	pending := []Pulse{Pulse{
		origin: button,
		target: startNode,
		high:   false,
	}}
	for len(pending) > 0 {
		curr := pending[0]
		pending = pending[1:]
		if curr.high {
			highPulses += 1
		} else {
			lowPulses += 1
		}
		if node, ok := m.Nodes[curr.target]; ok {
			pending = append(pending, node.Trigger(curr)...)
		}
	}
	return highPulses, lowPulses
}

func (m *Machine) GetState() int {
	state := 0
	var names []string
	for name := range m.Nodes {
		names = append(names, name)
	}
	slices.Sort(names)
	for _, name := range names {
		flipflop, ok := m.Nodes[name].(*FlipFlop)
		if !ok {
			continue
		}
		state <<= 1
		if flipflop.State {
			state |= 1
		}
	}
	return state
}
