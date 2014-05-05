package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func primeNumbers(n int) chan int {
	s := make([]bool, n)

	for x := 2; x < int(math.Sqrt(float64(n)))+1; x++ {
		if !s[x] {
			for i := x + x; i < len(s); i = i + x {
				s[i] = true
			}
		}
	}

	ch := make(chan int)
	go func() {
		for i := 2; i < n; i++ {
			if !s[i] {
				ch <- i
			}
		}
		close(ch)

	}()
	return ch
}

func main() {
	flag.Parse()

	read_file, _ := os.OpenFile(flag.Arg(0), os.O_RDONLY, 0600)
	defer read_file.Close()
	reader := bufio.NewReader(read_file)

	for {
		switch line, err := reader.ReadString('\n'); err {
		case nil:
			chopped := line[0 : len(line)-1]
			n, _ := strconv.Atoi(chopped)
			result := primeNumbers(n)

			response := make([]string, 0)
			for p := range result {
				response = append(response, strconv.Itoa(p))
			}
			fmt.Fprintln(os.Stdout, strings.Join(response, ","))

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
