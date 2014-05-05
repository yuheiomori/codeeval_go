package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func MultipleOfANumber(x, n int) (result int) {

	ch := make(chan int)
	original_n := n

	go func() {
		for {
			n = n + original_n
			ch <- n
		}
	}()

	for i := range ch {
		if i >= x {
			result = i
			break
		}
	}
	return
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
			integers := strings.Split(chopped, ",")
			x, _ := strconv.Atoi(integers[0])
			n, _ := strconv.Atoi(integers[1])
			result := MultipleOfANumber(x, n)
			fmt.Fprintln(os.Stdout, result)

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
