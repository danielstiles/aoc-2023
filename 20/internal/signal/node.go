package signal

type Pulse struct {
	origin string
	target string
	high   bool
}

type Node interface {
	Trigger(Pulse) []Pulse
	AddInput(string)
	AddTarget(string)
	GetName() string
	GetTargets() []string
}

type baseNode struct {
	Name    string
	Targets []string
	pulses  []Pulse
}

func (n *baseNode) AddInput(input string) {}

func (n *baseNode) AddTarget(target string) {
	n.Targets = append(n.Targets, target)
}

func (n *baseNode) GetName() string {
	return n.Name
}

func (n *baseNode) GetTargets() []string {
	return n.Targets
}

func (n *baseNode) initPulses() {
	n.pulses = make([]Pulse, len(n.Targets))
	for i, target := range n.Targets {
		n.pulses[i] = Pulse{
			origin: n.Name,
			target: target,
		}
	}
}

type FlipFlop struct {
	State bool
	baseNode
}

func (n *FlipFlop) Trigger(pulse Pulse) []Pulse {
	if pulse.high {
		return nil
	}
	n.State = !n.State
	if n.pulses == nil {
		n.initPulses()
	}
	for i := range n.Targets {
		n.pulses[i].high = n.State
	}
	return n.pulses
}

type Conjunction struct {
	State map[string]bool
	baseNode
}

func (n *Conjunction) Trigger(pulse Pulse) []Pulse {
	n.State[pulse.origin] = pulse.high
	allHigh := true
	for _, state := range n.State {
		allHigh = allHigh && state
	}
	if n.pulses == nil {
		n.initPulses()
	}
	for i := range n.Targets {
		n.pulses[i].high = !allHigh
	}
	return n.pulses
}

func (n *Conjunction) AddInput(input string) {
	if n.State == nil {
		n.State = make(map[string]bool)
	}
	n.State[input] = false
}

type Broadcast struct {
	baseNode
}

func (n *Broadcast) Trigger(pulse Pulse) []Pulse {
	if n.pulses == nil {
		n.initPulses()
	}
	for i := range n.Targets {
		n.pulses[i].high = pulse.high
	}
	return n.pulses
}
