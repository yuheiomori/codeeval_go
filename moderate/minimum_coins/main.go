package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var COINS_VALUES = []int{5, 3, 1}

func solve(s string) int {

	value, _ := strconv.Atoi(s)
	total_coin_count := 0

	for _, coin_value := range COINS_VALUES {
		coin_count := value / coin_value
		total_coin_count += coin_count
		value -= coin_count * coin_value
		if value == 0 {
			break
		}

	}

	return total_coin_count
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
