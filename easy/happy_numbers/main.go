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

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IsHappyNumber(s string) (result int) {

	mem := make([]int, 0)

	var _inner func(_s string, _m []int) (_r int)
	_inner = func(s string, mem []int) (result int) {
		char_list := strings.Split(s, "")

		candidate := 0
		for i := 0; i < len(char_list); i++ {
			n, _ := strconv.Atoi(char_list[i])
			candidate += n * n
		}

		if candidate == 1 {
			result = 1
		} else if intInSlice(candidate, mem) {
			result = 0
		} else {
			mem = append(mem, candidate)
			result = _inner(strconv.Itoa(candidate), mem)
		}
		return
	}

	return _inner(s, mem)

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
			result := IsHappyNumber(chopped)
			fmt.Fprintln(os.Stdout, result)
		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
