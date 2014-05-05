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

type StringArray []string

func (sa *StringArray) IntArray() []int {
	result := make([]int, len(*sa))
	for idx, c := range *sa {
		i, _ := strconv.Atoi(c)
		result[idx] = i
	}
	return result
}

func isJolly(l []int) bool {
	n := len(l) - 1
	memo := make(map[int]bool)
	for i := 1; i <= n; i++ {
		diff := int(math.Abs(float64(l[i] - l[i-1])))
		if diff < 1 || diff > n || memo[diff] {
			return false
		}
		memo[diff] = true
	}
	return true
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
			var l StringArray = strings.Split(chopped, " ")
			result := isJolly(l.IntArray()[1:])

			if result {
				fmt.Fprintln(os.Stdout, "Jolly")
			} else {
				fmt.Fprintln(os.Stdout, "Not jolly")
			}

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
