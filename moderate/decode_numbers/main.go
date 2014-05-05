package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func solve(s string) int {
	if len(s) <= 1 {
		return 1
	}

	count := 0
	char_length := 1

	for {
		if len(s) < char_length {
			break
		}

		target := s[:char_length]
		if len(target) != char_length {
			break
		}
		if i, _ := strconv.Atoi(target); i > 26 {
			break
		}
		count += solve(s[char_length:])
		char_length += 1
	}

	return count
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
