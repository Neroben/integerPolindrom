package main

import (
	"fmt"
	"os"
)

func main() {
	var number = input()
	fmt.Printf("isPalindrome = %t\n", isPalindrome(number))
}

func input() int {
	fmt.Printf("Введите число для проверки: ")
	var number int
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		os.Exit(404)
	}
	return number
}

func isPalindrome(number int) bool {
	if number == reverseInteger(number) {
		return true
	}
	return false
}

func reverseInteger(number int) int {
	reverseNumber := 0
	for number != 0 {
		reverseNumber += number % 10
		reverseNumber *= 10
		number /= 10
	}
	return reverseNumber / 10
}
