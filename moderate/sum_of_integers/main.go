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

func max(a, b int) (s int) {
	if a > b {
		return a
	}
	return b
}

func sum_of_integers(ints []int) int {
	var max_ending_here int
	var max_so_far int = math.MinInt32
	var max_element int = math.MinInt32
	for _, i := range ints {
		max_ending_here = max(0, max_ending_here+i)
		max_so_far = max(max_ending_here, max_so_far)
		max_element = max(max_element, i)
	}
	if max_so_far != 0 {
		return max_so_far
	}
	return max_element

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
			list := strings.Split(chopped, ",")
			integers := []int{}
			for _, c := range list {
				stripped := strings.Replace(string(c), " ", "", -1)
				i, _ := strconv.Atoi(stripped)
				integers = append(integers, i)
			}
			result := sum_of_integers(integers)
			fmt.Fprintln(os.Stdout, result)

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
