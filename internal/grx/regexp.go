package grx

import (
	"sort"
	"strconv"
)

//CreateRegExp CreateRegExp
func CreateRegExp(data []string) string {
	if len(data) < 2 {
		return ""
	}
	dataWidth := -1
	for _, d := range data {
		if dataWidth == -1 {
			dataWidth = len(d)
		} else if len(d) != dataWidth {
			return ""
		}
	}
	sort.Sort(stringSlice(data))
	minValue, _ := strconv.ParseInt(data[0], 10, 32)
	maxValue, _ := strconv.ParseInt(data[len(data)-1], 10, 32)

	arx := make([][][]byte, 0)

	for i := 0; i < len(data); i++ {
		p := []byte(data[i])
		rx := make([][]byte, 0)
		for x := 0; x < dataWidth; x++ {
			rx = append(rx, make([]byte, 0))
		}
		for x := 0; x < dataWidth; x++ {
			rx[x] = append(rx[x], p[x])
		}
		arx = append(arx, rx)
	}

	for i := 0; i < dataWidth; i++ {
		arx = calculateRegexp(arx)
	}
	str := toString(arx)

	if !validate(data, str, minValue, maxValue) {
		return ""
	}

	return str
}

func calculateRegexp(arx [][][]byte) [][][]byte {
	ret := make([][][]byte, 0)
	for i := 0; i < len(arx); i++ {
		p := arx[i]
		tmp := make([]byte, 0)
		pri := true
		base := 0
		for base = 0; base < len(p); base++ {
			if len(p[base]) > 1 {
				base--
				break
			}
		}
		if base == len(p) {
			base--
		}
		act := i

		for i++; i < len(arx); i++ {
			a := arx[i]
			ca, cb, to, eq := compareArray(p, a)
			_ = cb
			if ca-to == 1 && eq {
				if pri {
					tmp = append(tmp, p[to][0])
					pri = false
				}

				tmp = append(tmp, a[to][0])
				arx[i] = make([][]byte, 0)
			} else {
				i--
				break
			}
		}
		if len(tmp) > 0 {
			sort.Sort(byteSlice(tmp))
			arx[act][base] = tmp
		}
	}
	for i := 0; i < len(arx); i++ {
		if len(arx[i]) == 0 {
			continue
		}
		ret = append(ret, arx[i])
	}
	return ret
}
