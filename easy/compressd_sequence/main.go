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

func CompressSequence(s []string) []string {

	array := make([]string, 0)
	last := s[0]
	count := 0
	for _, r := range s {
		if r == last {
			count++
		} else {
			array = append(array, strconv.Itoa(count))
			array = append(array, last)
			last = r
			count = 1
		}
	}
	array = append(array, strconv.Itoa(count))
	array = append(array, last)

	return array

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
			result := CompressSequence(strings.Split(chopped, " "))
			fmt.Fprintln(os.Stdout, strings.Join(result, " "))

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
