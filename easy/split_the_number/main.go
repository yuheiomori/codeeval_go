package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	calc_pattern = regexp.MustCompile("([a-z]+)([+-])[a-z]+")
)

func solve(s string) (result string, err error) {

	elements := strings.Split(s, " ")
	digit := elements[0]
	pattern := elements[1]

	submatches := calc_pattern.FindStringSubmatch(pattern)

	lp := submatches[1]
	o := submatches[2]

	l, _ := strconv.Atoi(digit[:len(lp)])
	r, _ := strconv.Atoi(digit[len(lp):])

	if o == "+" {
		result = strconv.Itoa(l + r)
	} else if o == "-" {
		result = strconv.Itoa(l - r)
	} else {
		err = errors.New("argument error")
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
			if len(chopped) == 0 {
				continue
			}

			result, err := solve(chopped)
			if err == nil {
				fmt.Fprintln(os.Stdout, result)
			}

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
