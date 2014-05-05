package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func solve(l []string) (result string) {
	re := regexp.MustCompile("\\d+")

	digits := make([]string, 0)
	words := make([]string, 0)

	for _, word := range l {
		if re.Match([]byte(word)) {
			digits = append(digits, word)
		} else {
			words = append(words, word)
		}
	}

	result = strings.Join(words, ",")
	if len(words) > 0 && len(digits) > 0 {
		result = result + "|"
	}
	result = result + strings.Join(digits, ",")
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
			result := solve(strings.Split(chopped, ","))
			fmt.Fprintln(os.Stdout, result)

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
