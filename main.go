package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/coc1961/goregexp/internal/grx"
)

func main() {

	start := false

	go func(start *bool) {

		<-time.After(time.Second * 3)
		if !*start {
			fmt.Println("I can't read the Numbers List from stdin")
			os.Exit(1)
		}

	}(&start)

	data := make([]string, 0)
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _, err := reader.ReadLine()
		if !start {
			start = true
		}
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
