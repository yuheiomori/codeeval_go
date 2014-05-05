package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func pangrams(s string) (result string) {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	trimed := strings.Replace(s, " ", "", -1)
	lower := strings.ToLower(trimed)

	for _, c := range lower {
		alphabet = strings.Replace(alphabet, string(c), "", -1)
	}

	if len(alphabet) == 0 {
		result = "NULL"
	} else {
		result = alphabet
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
			result := pangrams(chopped)
			fmt.Fprintln(os.Stdout, result)

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
