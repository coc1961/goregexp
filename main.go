package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/coc1961/goregexp/internal/grx"
)

func main() {

	data := make([]string, 0)
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		str := strings.Trim(string(text), "\n")
		if str != "" {
			data = append(data, str)
		}
	}

	reg := grx.CreateRegExp(data)
	if reg == "" {
		return
	}
	fmt.Println(reg)

}
