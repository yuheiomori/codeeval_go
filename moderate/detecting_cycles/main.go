package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type stringSlice []string

func (slice stringSlice) index(value string) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}

func (slice stringSlice) equals(others []string) bool {
	if len(slice) != len(others) {
		return false
	}
	for i, v := range others {
		if v != slice[i] {
			return false
		}
	}
	return true
}

func (slice stringSlice) push(value string) stringSlice {
	f := stringSlice{value}
	return append(f, slice...)
}

func DetectingCycles(l stringSlice) (result []string) {

	defer func() {
		// break silently if index out of bound error occured.
		recover()
	}()

	for {
		top := l[0]
		l = l[1:]
		pos := l.index(top)
		if pos == -1 {
			continue
		}
		candidate := l[0:pos]
		sample := l[pos+1 : pos+1+len(candidate)]
		if candidate.equals(sample) {
			return candidate.push(top)

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
			integers := strings.Split(chopped, " ")
			result := DetectingCycles(integers)
			fmt.Fprintln(os.Stdout, strings.Join(result, " "))

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
