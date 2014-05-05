package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func solve(numbers []string) (r int) {
	m := make(map[string]int)

	// count occurance
	for _, n := range numbers {
		m[n] += 1
	}

	// pickup candidates
	candidates := make([]string, 0)
	for k, v := range m {
		if v == 1 {
			candidates = append(candidates, k)
		}
	}

	// determin pos
	if len(candidates) != 0 {
		sort.Strings(candidates)
		lowest_common_number := candidates[0]

		for idx, n := range numbers {
			if n == lowest_common_number {
				r = idx + 1
			}
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
			result := solve(strings.Split(chopped, " "))
			fmt.Fprintln(os.Stdout, result)

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
