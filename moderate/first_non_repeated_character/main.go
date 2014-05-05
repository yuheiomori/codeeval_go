package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func Map(chars []string, f func(x string) int) []int {

	result := make([]int, len(chars))
	for i, c := range chars {
		result[i] = f(c)
	}
	return result
}

func Index(ints []int, t int) int {
	for idx, i := range ints {
		if i == t {
			return idx
		}
	}
	return -1
}

func firstNonRepeatedCharacter(s string) string {
	chars := strings.Split(s, "")
	counts := Map(chars, func(x string) int {
		return strings.Count(s, x)
	})

	return chars[Index(counts, 1)]

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
			result := firstNonRepeatedCharacter(chopped)
			fmt.Fprintln(os.Stdout, result)

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
