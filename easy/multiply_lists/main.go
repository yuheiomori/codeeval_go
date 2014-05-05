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

func solve(s string) string {
	arrays := strings.Split(s, "|")
	array1 := strings.Split(strings.Trim(arrays[0], " "), " ")
	array2 := strings.Split(strings.Trim(arrays[1], " "), " ")

	min_length := int(math.Min(float64(len(array1)), float64(len(array2))))
	results := make([]string, min_length)

	for i := 0; i < min_length; i++ {
		a, _ := strconv.Atoi(array1[i])
		b, _ := strconv.Atoi(array2[i])
		results[i] = strconv.Itoa(a * b)
	}

	return strings.Join(results, " ")
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
