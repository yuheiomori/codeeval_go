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

func mod(n, m int) (result int) {
	if n < m {
		return n
	}
	return mod(n-m, m)
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
			n, _ := strconv.Atoi(integers[0])
			m, _ := strconv.Atoi(integers[1])
			result := mod(n, m)
			fmt.Fprintln(os.Stdout, result)
		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
