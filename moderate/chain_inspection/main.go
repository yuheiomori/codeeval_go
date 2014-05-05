package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func solve(s string) string {

	d := map[string]string{}
	elements := strings.Split(s, ";")

	for _, e := range elements {
		splitted := strings.Split(e, "-")
		f, t := splitted[0], splitted[1]
		d[f] = t
	}

	passed_item := map[string]bool{}
	elem := "BEGIN"
	for {

		if _, ok := passed_item[elem]; ok {
			return "BAD"

		} else if elem == "END" {
			break

		} else if _, ok := d[elem]; !ok {
			return "BAD"

		}
		passed_item[elem] = true
		elem = d[elem]
	}

	if len(d) == len(passed_item) {
		return "GOOD"
	} else {
		return "BAD"
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

			result := solve(chopped)
			fmt.Fprintln(os.Stdout, result)

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
