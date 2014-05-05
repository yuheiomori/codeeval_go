package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func removeCharacters(target, remove_string string) string {
	chars2 := make([]string, 0)
	for _, rune := range target {
		c := string(rune)

		if c == " " || strings.Index(remove_string, c) == -1 {
			chars2 = append(chars2, c)
		}
	}

	return strings.Join(chars2, "")
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
			columns := strings.Split(chopped, ",")
			result := removeCharacters(columns[0], columns[1])
			fmt.Fprintln(os.Stdout, result)

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
