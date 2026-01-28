package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err == nil {
			sum += num
		}
	}
	if err := scanner.Err(); err != nil && err != io.EOF {
		fmt.Fprintln(os.Stderr, "Ошибка:", err)
		return
	}
	fmt.Println(sum)
	//Чтобы вывелась сумма, надо нажать Ctrl+Z
}
