package main

import (
	"fmt"
)

func fn(num uint) uint {
	result := uint(0)
	multiplier := uint(1)
	for num > 0 {
		digit := num % 10
		if digit != 0 && digit%2 == 0 {
			result += digit * multiplier
			multiplier *= 10
		}
		num /= 10
	}
	if result == 0 {
		return 100
	}
	return result
}

func main() {
	var input uint
	fmt.Scan(&input)
	result := fn(input)
	fmt.Printf("%d\n", result)
}
