package grx

import (
	"fmt"
	"sort"
	"strings"
)

type printByte [][]byte

func (b printByte) String() string {
	str := ""
	for _, a := range b {
		if len(a) == 1 {
			str += string(a[0])
		} else {
			str += "["
			coma := ""
			sort.Sort(byteSlice(a))

			x := byte(255)
			x1 := byte(255)
			for _, a1 := range a {
				if x == byte(255) {
					x = a1
					x1 = a1
					continue
				}
				if x1+1 == a1 {
					x1 = a1
				} else {
					if x == x1 {
						str += coma + fmt.Sprintf("%s", string(x))
					} else {
						str += coma + fmt.Sprintf("%s-%s", string(x), string(x1))
					}
					coma = ","
					x = a1
					x1 = a1
				}
			}
			if x == x1 {
				str += coma + fmt.Sprintf("%s", string(x))
			} else {
				str += coma + fmt.Sprintf("%s-%s", string(x), string(x1))
			}
			str += "]"
		}
	}

	return str
}

func toString(arx [][][]byte) string {
	str := "^("
	for i := 0; i < len(arx); i++ {
		if len(arx[i]) == 0 {
			continue
		}
		str += fmt.Sprintf("%s|", printByte(arx[i]))
	}
	str = strings.TrimRight(str, "|")
	str += ")$"
	return str
}
