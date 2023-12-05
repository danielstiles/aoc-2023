package calibration

func GetCalibration(line []byte) int {
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
