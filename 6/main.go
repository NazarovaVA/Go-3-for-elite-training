package main

import (
	"fmt"
	"strings"
)

type Battery struct {
	state string
}

func (b Battery) String() string {
	var chars []string
	for _, char := range b.state {
		if char == '1' {
			chars = append(chars, "X")
		} else {
			chars = append(chars, " ")
		}
	}
	result := "["
	xCount := strings.Count(b.state, "1")
	result += strings.Repeat(" ", len(b.state)-xCount)
	result += strings.Repeat("X", xCount)
	return result + "]"
}

func main() {
	var input string
	fmt.Scan(&input)
	battery := Battery{
		state: input,
	}
	fmt.Println(battery)
}
