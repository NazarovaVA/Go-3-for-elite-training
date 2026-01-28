package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Ошибка чтения:", err)
		return
	}
	line = strings.TrimSuffix(line, "\r\n")
	numbers := strings.Split(line, ";")
	num1Str := strings.ReplaceAll(numbers[0], " ", "")
	num1Str = strings.Replace(num1Str, ",", ".", 1)
	num2Str := strings.ReplaceAll(numbers[1], " ", "")
	num2Str = strings.Replace(num2Str, ",", ".", 1)
	num1, err := strconv.ParseFloat(num1Str, 64)
	if err != nil {
		fmt.Println("Ошибка преобразования первого числа:", err)
		return
	}
	num2, err := strconv.ParseFloat(num2Str, 64)
	if err != nil {
		fmt.Println("Ошибка преобразования второго числа:", err)
		return
	}
	if num2 == 0 {
		fmt.Println("Ошибка: деление на ноль")
		return
	}
	result := num1 / num2
	fmt.Printf("%.4f\n", result)
}
