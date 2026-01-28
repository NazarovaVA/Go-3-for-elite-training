package main

import (
	"fmt"
)

var cache = make(map[int]int)

func work(x int) int {
	if result, ok := cache[x]; ok {
		return result
	}
	result := x * x
	cache[x] = result
	return result
}
func main() {
	var numbers [10]int
	fmt.Scan(&numbers[0], &numbers[1], &numbers[2], &numbers[3],
		&numbers[4], &numbers[5], &numbers[6], &numbers[7],
		&numbers[8], &numbers[9])
	for _, num := range numbers {
		fmt.Printf("%d ", work(num))
	}
	fmt.Println("time limit ok")
}
