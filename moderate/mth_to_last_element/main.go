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

func printMthLastElement(elements []string) {

	list := elements[0 : len(elements)-1]
	m, _ := strconv.Atoi(elements[len(elements)-1])
	if m > 0 && m <= len(list) {
		fmt.Fprintln(os.Stdout, list[len(list)-m])
	}
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
			elements := strings.Split(chopped, " ")
			printMthLastElement(elements)

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
