package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode"
)

var (
	m = map[rune]rune{
		'a': '0',
		'b': '1',
		'c': '2',
		'd': '3',
		'e': '4',
		'f': '5',
		'g': '6',
		'h': '7',
		'i': '8',
		'j': '9',
	}
)

func findHiddenDigits(s string) (result string) {
	tmp := make([]rune, 0)

	for _, r := range s {
		if unicode.IsDigit(r) {
			tmp = append(tmp, r)
			continue
		}

		if d, ok := m[r]; ok {
			tmp = append(tmp, d)
			continue
		}
	}
	result = string(tmp)
	if result == "" {
		result = "NONE"
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
			result := findHiddenDigits(chopped)
			fmt.Fprintln(os.Stdout, result)

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
