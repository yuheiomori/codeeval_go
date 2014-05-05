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

var (
	SYMBOLS      = []string{"I", "X", "C", "M"}
	HALF_SYMBOLS = []string{"V", "L", "D"}
)

func solve(s string) (result string) {

	length := len(s)

	for idx, r := range s {
		digit, _ := strconv.Atoi(string(r))
		if digit == 0 {
			continue
		}
		order := length - idx

		if digit == 4 {
			result += SYMBOLS[order-1] + HALF_SYMBOLS[order-1]
		} else if digit == 9 {
			result += SYMBOLS[order-1] + SYMBOLS[order]

		} else if digit == 5 {
			result += HALF_SYMBOLS[order-1]
		} else if digit > 5 {
			result += HALF_SYMBOLS[order-1] + strings.Repeat(SYMBOLS[order-1], (digit-5))
		} else {
			result += strings.Repeat(SYMBOLS[order-1], digit)
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

			result := solve(chopped)
			fmt.Fprintln(os.Stdout, result)

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
