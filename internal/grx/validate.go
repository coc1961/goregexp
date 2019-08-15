package grx

import (
	"fmt"
	"regexp"
	"strings"
)

func validate(arr []string, reg string, min, max int64) bool {
	for _, b := range arr {
		matched, err := regexp.Match(reg, []byte(strings.Trim(b, " ")))
		if !matched {
			fmt.Println("No Match", b, matched, err)
			return false
		}

	}

	for i := min - 1000; i <= max+1000; i++ {
		str := fmt.Sprintf("%d", i)
		esta := false
		for _, b := range arr {
			if str == strings.Trim(b, " ") {
				esta = true
			}
		}
		if !esta {
			matched, err := regexp.Match(reg, []byte(strings.Trim(str, " ")))
			if matched {
				fmt.Println("Match ERROR", str, err)
				return false
			}
		}
	}
	return true
}
