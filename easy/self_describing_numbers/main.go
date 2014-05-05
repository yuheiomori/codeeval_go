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

func IsSelfDescribingNumber(s string) (result int) {

	char_list := strings.Split(s, "")
	result = 1

	for i, c := range char_list {
		num, _ := strconv.Atoi(c)
		idx := strconv.Itoa(i)

		if !(strings.Count(s, idx) == num) {
			result = 0
			break
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
			result := IsSelfDescribingNumber(chopped)
			fmt.Fprintln(os.Stdout, result)
		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
