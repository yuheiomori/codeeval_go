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

func solve(s, key string) string {
	postions := strings.Split(key, " ")
	runes := []rune(s)
	result_runes := make([]rune, len(postions))

	for idx, k := range postions {
		pos, _ := strconv.Atoi(k)
		if pos-1 < len(runes) {
			result_runes[idx] = runes[pos-1]
		}
	}
	return string(result_runes)
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
			if chopped == "" {
				continue
			}
			tmp := strings.Split(chopped, "| ")
			result := solve(tmp[0], tmp[1])
			fmt.Fprintln(os.Stdout, result)

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
