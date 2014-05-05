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

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPandronome(n int) bool {
	t1 := strconv.Itoa(n)
	t2 := Reverse(t1)
	return t1 == t2
}

func reverseAndAdd(n int) (c, p int) {

	var inner_func func(nn int) int
	inner_func = func(nn int) int {
		r, _ := strconv.Atoi(Reverse(strconv.Itoa(nn)))
		result := nn + r
		c++

		if isPandronome(result) {
			return result
		}
		return inner_func(result)
	}

	return c, inner_func(n)
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
			c, p := reverseAndAdd(n)
			fmt.Fprintln(os.Stdout, strings.Join([]string{
				strconv.Itoa(c),
				strconv.Itoa(p)},
				" "))

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
