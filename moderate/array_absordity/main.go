package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func solve(s string) (result []string) {

	splitted := strings.Split(s, ";")
	elements := strings.Split(splitted[1], ",")
	m := make(map[string]int)

	for _, e := range elements {

		if m[e] == 1 {
			result = append(result, e)
		} else if m[e] == 0 {
			m[e]++
		}
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
			result := solve(chopped)
			fmt.Fprintln(os.Stdout, strings.Join(result, ","))

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
