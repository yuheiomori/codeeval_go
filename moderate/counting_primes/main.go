package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool {
	for i := 2; i < n/2+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs := strings.Split(scanner.Text(), ",")
		from, _ := strconv.Atoi(inputs[0])
		to, _ := strconv.Atoi(inputs[1])

		cnt := 0
		for i := from; i <= to; i++ {
			if isPrime(i) {
				cnt++
			}
		}

		fmt.Println(cnt)
	}
}
