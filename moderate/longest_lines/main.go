package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Row struct {
	Line   string
	Length int
}

type Lines []Row

func (l Lines) Len() int {
	return len(l)
}

func (l Lines) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l Lines) Less(i, j int) bool {
	return l[i].Length > l[j].Length
}

func main() {
	flag.Parse()

	read_file, _ := os.OpenFile(flag.Arg(0), os.O_RDONLY, 0600)
	defer read_file.Close()
	reader := bufio.NewReader(read_file)

	switch content, err := ioutil.ReadAll(reader); err {

	case nil:
		var lines Lines

		content_list := strings.Split(string(content), "\n")
		n, _ := strconv.Atoi(content_list[0])

		for _, s := range content_list[1:] {
			lines = append(lines, Row{s, len(s)})
		}

		sort.Sort(lines)

		for i := 0; i < n; i++ {
			fmt.Fprintln(os.Stdout, lines[i].Line)
		}

	case io.EOF:
		os.Exit(0)
	default:
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)

	}
}
