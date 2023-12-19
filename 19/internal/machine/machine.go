package machine

import (
	"regexp"
	"strconv"
)

var numExpr = regexp.MustCompile("\\d+")
var condExpr = regexp.MustCompile("(x|s|a|m)(<|>)(\\d+):(\\w+)")
var ruleExpr = regexp.MustCompile("(\\w+){.*,(\\w+)}")

type Part struct {
	X int
	M int
	A int
	S int
}

func ParsePart(line []byte) Part {
	nums := numExpr.FindAll(line, -1)
	x, _ := strconv.Atoi(string(nums[0]))
	m, _ := strconv.Atoi(string(nums[1]))
	a, _ := strconv.Atoi(string(nums[2]))
	s, _ := strconv.Atoi(string(nums[3]))
	return Part{
		X: x,
		M: m,
		A: a,
		S: s,
	}
}

const (
	X            = byte('x')
	M            = byte('m')
	A            = byte('a')
	S            = byte('s')
	Greater      = byte('>')
	GreaterEqual = byte('.')
	Less         = byte('<')
	LessEqual    = byte(',')
)

type Variable byte

type Cond struct {
	Variable    byte
	Condition   byte
	Level       int
	Origin      string
	Destination string
}

func (c Cond) Invert() Cond {
	var newCondition byte
	switch c.Condition {
	case Greater:
		newCondition = LessEqual
	case Less:
		newCondition = GreaterEqual
	case GreaterEqual:
		newCondition = Less
	case LessEqual:
		newCondition = Greater
	}
	return Cond{
		Variable:  c.Variable,
		Condition: newCondition,
		Level:     c.Level,
	}
}

type Rule struct {
	Name       string
	Conditions []Cond
	Default    string
}

var Accept = &Rule{Name: "A"}

var Reject = &Rule{Name: "R"}

func ParseRule(line []byte) *Rule {
	ruleMatch := ruleExpr.FindSubmatch(line)
	condMatch := condExpr.FindAllSubmatch(line, -1)
	var conds []Cond
	ruleName := string(ruleMatch[1])
	for _, cond := range condMatch {
		level, _ := strconv.Atoi(string(cond[3]))
		conds = append(conds, Cond{
			Variable:    cond[1][0],
			Condition:   cond[2][0],
			Level:       level,
			Origin:      ruleName,
			Destination: string(cond[4]),
		})
	}
	return &Rule{
		Name:       ruleName,
		Conditions: conds,
		Default:    string(ruleMatch[2]),
	}
}

func (r *Rule) Check(part Part) string {
	for _, cond := range r.Conditions {
		var val int
		switch cond.Variable {
		case X:
			val = part.X
		case M:
			val = part.M
		case A:
			val = part.A
		case S:
			val = part.S
		}
		switch cond.Condition {
		case Greater:
			if val > cond.Level {
				return cond.Destination
			}
		case Less:
			if val < cond.Level {
				return cond.Destination
			}
		}
	}
	return r.Default
}

const (
	StartRule = "in"
)

type Program struct {
	Rules map[string]*Rule
}

func NewProgram() *Program {
	p := &Program{
		Rules: make(map[string]*Rule),
	}
	p.AddRule(Accept)
	p.AddRule(Reject)
	return p
}

func (p *Program) AddRule(r *Rule) {
	p.Rules[r.Name] = r
}

func (p *Program) Process(part Part) bool {
	dest := p.Rules[StartRule].Check(part)
	for p.Rules[dest] != Accept && p.Rules[dest] != Reject {
		dest = p.Rules[dest].Check(part)
	}
	return p.Rules[dest] == Accept
}

func (p *Program) GetAccepatableRanges() []*Rule {
	finished := true
	ranges := p.reaches(Accept)
	for _, rule := range ranges {
		if rule.Name != StartRule {
			finished = false
		}
	}
	for !finished {
		finished = true
		var newRanges []*Rule
		for _, rule := range ranges {
			if rule.Name == StartRule {
				newRanges = append(newRanges, rule)
				continue
			}
			finished = false
			reach := p.reaches(p.Rules[rule.Name])
			for _, reachRule := range reach {
				if len(reachRule.Conditions) > 0 {
					for _, cond := range rule.Conditions {
						if cond == reachRule.Conditions[0] {
							continue
						}
					}
				}
				newRanges = append(newRanges, &Rule{
					Name:       reachRule.Name,
					Conditions: append(reachRule.Conditions, rule.Conditions...),
				})
			}
		}
		ranges = newRanges
	}
	return ranges
}

func (p *Program) reaches(dest *Rule) []*Rule {
	var reach []*Rule
	for _, r := range p.Rules {
		var mustFail []Cond
		for _, cond := range r.Conditions {
			if p.Rules[cond.Destination] == dest {
				reach = append(reach, &Rule{
					Name:       r.Name,
					Conditions: append(mustFail, cond),
				})
			}
			mustFail = append(mustFail, cond.Invert())
		}
		if p.Rules[r.Default] == dest {
			reach = append(reach, &Rule{
				Name:       r.Name,
				Conditions: mustFail,
			})
		}
	}
	return reach
}
