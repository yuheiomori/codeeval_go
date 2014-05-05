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

func BitPositions(x, p1, p2 int) (result bool) {
	b := strconv.FormatInt(int64(x), 2)
	return b[len(b)-p1] == b[len(b)-p2]
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
			x, _ := strconv.Atoi(integers[0])
			p1, _ := strconv.Atoi(integers[1])
			p2, _ := strconv.Atoi(integers[2])
			result := BitPositions(x, p1, p2)
			fmt.Fprintln(os.Stdout, result)

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
