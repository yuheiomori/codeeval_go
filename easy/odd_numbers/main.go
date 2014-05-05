package main

import (
	"fmt"
	"os"
)

func OddGenerator(ch chan<- int) {
	for i := 1; ; i++ {
		if i%2 != 0 {
			ch <- i
		}

	}
}

func main() {
	ch := make(chan int)
	defer close(ch)
	go OddGenerator(ch)

	for odd_number := range ch {
		if odd_number > 100 {
			os.Exit(0)
		}
		fmt.Println(odd_number)
	}

}
