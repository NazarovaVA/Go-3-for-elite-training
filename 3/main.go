package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func adding(str1, str2 string) int64 {
	cleanString := func(s string) string {
		var result strings.Builder
		for _, r := range s {
			if unicode.IsDigit(r) {
				result.WriteRune(r)
			}
		}
		return result.String()
	}
	num1Str := cleanString(str1)
	num2Str := cleanString(str2)
	num1, _ := strconv.ParseInt(num1Str, 10, 64)
	num2, _ := strconv.ParseInt(num2Str, 10, 64)
	return num1 + num2
}

func main() {
	var input1, input2 string
	fmt.Scanln(&input1)
	fmt.Scanln(&input2)
	result := adding(input1, input2)
	fmt.Printf("%d", result)
}
