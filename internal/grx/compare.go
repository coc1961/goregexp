package grx

func compareBase(a, b [][]byte) (int, int, int, bool) {
	if len(a) != len(b) {
		return -1, -1, -1, false
	}
	conta := 0
	contb := 0
	tot := 0

	for i := 0; i < len(a); i++ {
		if len(a[i]) == 1 {
			conta++
		} else {
			break
		}
	}
	for i := 0; i < len(b); i++ {
		if len(b[i]) == 1 {
			contb++
		} else {
			break
		}
	}

	if conta != contb {
		return -1, -1, -1, false
	}

	for i := 0; i < conta; i++ {
		if a[i][0] == b[i][0] {
			tot++
		} else {
			break
		}
	}
	ret := true
	for i := conta; i < len(a); i++ {
		if !compareArr(a[i], b[i]) {
			ret = false
		}
	}
	return conta, contb, tot, ret
}

func compareArr(b1, b2 []byte) bool {
	if len(b1) != len(b2) {
		return false
	}
	cont := 0
	for i := 0; i < len(b1); i++ {
		if b1[i] == b2[i] {
			cont++
		} else {
			break
		}
	}
	return cont == len(b1)
}

func contain(arr []byte, b byte) bool {
	for _, b1 := range arr {
		if b1 == b {
			return true
		}
	}
	return false
}
func compare(b1, b2 []byte) int {
	l := min(len(b1), len(b2))
	cont := 0
	for i := 0; i < l; i++ {
		if b1[i] == b2[i] {
			cont++
		} else {
			break
		}
	}
	return cont
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
