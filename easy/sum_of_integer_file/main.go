package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	flag.Parse()

	read_file, _ := os.OpenFile(flag.Arg(0), os.O_RDONLY, 0600)
	defer read_file.Close()
	reader := bufio.NewReader(read_file)

	result := 0

	for {
		switch line, err := reader.ReadString('\n'); err {
		case nil:
			d, _ := strconv.Atoi(line[0 : len(line)-1])
			result += d

		case io.EOF:
			fmt.Fprintln(os.Stdout, result)
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
