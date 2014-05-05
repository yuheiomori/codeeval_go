package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func Unique(list []string) []string {
	var last string
	i := 0
	for _, x := range list {
		if i == 0 || x != last {
			last = x
			list[i] = x
			i++
		}
	}
	return list[0:i]
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
			integers := strings.Split(chopped, ",")
			result := Unique(integers)
			fmt.Fprintln(os.Stdout, strings.Join(result, ","))
		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
