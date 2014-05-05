package main

import (
	"fmt"
	"os"
)

func IsPrime(n int) bool {

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func PrimeGenerator(ch chan<- int) {
	for i := 2; ; i++ {
		if IsPrime(i) {
			ch <- i
		}

	}
}

func main() {
	cnt := 0
	sum := 0
	ch := make(chan int)
	defer close(ch)

	go PrimeGenerator(ch)
	for prime := range ch {
		sum += prime
		cnt++

		if cnt == 1000 {
			fmt.Println(sum)
			os.Exit(0)
		}

	}
}
