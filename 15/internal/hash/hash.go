package hash

import (
	"regexp"
	"strconv"
)

var stepExpr = regexp.MustCompile("([^,-=]*)(-|=)(\\d+)?")

func HashLine(line []byte) int {
	total := 0
	steps := stepExpr.FindAll(line, -1)
	for _, step := range steps {
		total += Hash(step)
	}
	return total
}

func Hash(in []byte) int {
	hash := 0
	for _, c := range in {
		hash = AddChar(hash, c)
	}
	return hash
}

func AddChar(start int, char byte) int {
	return ((start + int(char)) * 17) % 256
}

type Lens struct {
	Label       string
	FocalLength int
}

const (
	Remove = byte('-')
	Add    = byte('=')
)

type HashMap map[int][]Lens

func NewHashMap(line []byte) HashMap {
	ret := HashMap(make(map[int][]Lens))
	steps := stepExpr.FindAllSubmatch(line, -1)
	for _, step := range steps {
		ret.Process(step)
	}
	return ret
}

func (m HashMap) Process(step [][]byte) {
	label := string(step[1])
	op := step[2][0]
	switch op {
	case Add:
		focalLength, _ := strconv.Atoi(string(step[3]))
		new := Lens{
			Label:       label,
			FocalLength: focalLength,
		}
		m.Add(new)
	case Remove:
		m.Remove(label)
	}
}

func (m HashMap) Find(label string) int {
	hash := Hash([]byte(label))
	var i int
	for i = 0; i < len(m[hash]); i++ {
		if m[hash][i].Label == label {
			break
		}
	}
	if i == len(m[hash]) {
		return -1
	}
	return i
}

func (m HashMap) Add(new Lens) {
	i := m.Find(new.Label)
	hash := Hash([]byte(new.Label))
	if i == -1 {
		m[hash] = append(m[hash], new)
		return
	}
	m[hash][i] = new
}

func (m HashMap) Remove(label string) {
	i := m.Find(label)
	if i == -1 {
		return
	}
	hash := Hash([]byte(label))
	for i = i + 1; i < len(m[hash]); i++ {
		m[hash][i-1] = m[hash][i]
	}
	m[hash] = m[hash][:len(m[hash])-1]
}

func (m HashMap) RollUp() int {
	total := 0
	for hash, lenses := range m {
		for slot, lens := range lenses {
			total += (1 + hash) * (1 + slot) * lens.FocalLength
		}
	}
	return total
}

func (m HashMap) Print() string {
	ret := ""
	for hash, box := range m {
		ret += strconv.Itoa(hash) + ": ["
		first := true
		for _, lens := range box {
			if !first {
				ret += ", "
			}
			ret += lens.Label + ":" + strconv.Itoa(lens.FocalLength)
			first = false
		}
		ret += "] "
	}
	return ret
}
