package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func copyResult(a []int) *[]int {
	b := make([]int, len(a))
	copy(b, a)
	return &b
}

func sum(a []int) int {
	s := 0
	for _, e := range a {
		s += e
	}
	return s
}

func combinations_gen(iterable []int, r int) chan *[]int {

	ch := make(chan *[]int)

	go func(ch chan *[]int) {
		defer close(ch)
		pool := iterable
		n := len(pool)

		if r > n {
			return
		}

		indices := make([]int, r)
		for i := range indices {
			indices[i] = i
		}

		result := make([]int, r)
		for i, el := range indices {
			result[i] = pool[el]
		}

		ch <- copyResult(result)

		for {
			i := r - 1
			for ; i >= 0 && indices[i] == i+n-r; i -= 1 {
			}

			if i < 0 {
				return
			}

			indices[i] += 1
			for j := i + 1; j < r; j += 1 {
				indices[j] = indices[j-1] + 1
			}

			for ; i < len(indices); i += 1 {
				result[i] = pool[indices[i]]
			}
			ch <- copyResult(result)
		}

	}(ch)
	return ch
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		digits := strings.Split(scanner.Text(), ",")
		intdigits := make([]int, len(digits))
		for i, v := range digits {
			intdigits[i], _ = strconv.Atoi(v)
		}

		cnt := 0
		for combination := range combinations_gen(intdigits, 4) {
			if sum(*combination) == 0 {
				cnt += 1
			}
		}
		fmt.Println(cnt)
	}
}
