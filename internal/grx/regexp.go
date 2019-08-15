package grx

import (
	"sort"
	"strconv"
)

//GeneroRegExp GeneroRegExp
func GeneroRegExp(data []string) string {
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
	sort.Sort(StringSlice(data))
	minv, _ := strconv.ParseInt(data[0], 10, 32)
	maxv, _ := strconv.ParseInt(data[len(data)-1], 10, 32)

	arx := make([][][]byte, 0)

	for i := 0; i < len(data); i++ {
		p := []byte(data[i])
		rx := make([][]byte, 0)
		for x := 0; x < dataWidth; x++ {
			rx = append(rx, make([]byte, 0))
		}
		diff := -1
		//fmt.Println(printByte1(p))
		for i++; i < len(data); i++ {
			a := []byte(data[i])
			if diff == -1 {
				diff = compare(p, a)
				for j := 0; j < len(p); j++ {
					rx[j] = append(rx[j], p[j])
				}
				if diff != len(p)-1 {
					i--
					break
				}
				//fmt.Println(diff)
			}
			if diff != compare(p, a) {
				i--
				break
			}
			//fmt.Println("  ", printByte1(a))
			for j := diff; j < min(len(p), len(a)); j++ {
				if !contain(rx[j], a[j]) {
					rx[j] = append(rx[j], a[j])
				}
			}
		}
		arx = append(arx, rx)
	}
	//fmt.Println(toString(arx))
	for i := 0; i < dataWidth; i++ {
		arx = compacto(arx)
	}
	str := toString(arx)

	if !validate(data, str, minv, maxv) {
		return ""
	}
	return str
}

func compacto(arx [][][]byte) [][][]byte {
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
		act := i

		for i++; i < len(arx); i++ {
			a := arx[i]
			ca, cb, to, eq := compareBase(p, a)
			_ = cb
			if ca-to == 1 && eq {
				//fmt.Println(ca, cb, to, eq)
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
			sort.Sort(ByteSlice(tmp))
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
