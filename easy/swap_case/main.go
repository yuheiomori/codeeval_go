package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode"
)

func solve(s string) string {
	runes := make([]rune, len(s))

	for idx, r := range []rune(s) {
		if !unicode.IsLetter(r) {
			runes[idx] = r
			continue
		}
		if unicode.IsUpper(r) {
			runes[idx] = unicode.ToLower(r)
		} else {
			runes[idx] = unicode.ToUpper(r)
		}

	}
	return string(runes)
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
