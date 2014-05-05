package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(i int) bool {
	s := strconv.Itoa(i)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func isPrime(i int) bool {
	for j := 2; j < i/2+1; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}

func main() {
	for i := 1000; i > 0; i-- {
		if isPrime(i) && isPalindrome(i) {
			fmt.Println(i)
			break
		}
	}

}
