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

func FizzBuzz(args []string) []string {
	a, _ := strconv.Atoi(args[0])
	b, _ := strconv.Atoi(args[1])
	c, _ := strconv.Atoi(args[2])

	rtnValue := make([]string, 0)

	for i := 1; i <= c; i++ {
		tmp := make([]string, 0)
		if i%a == 0 {
			tmp = append(tmp, "F")
		}
		if i%b == 0 {
			tmp = append(tmp, "B")

		}
		if len(tmp) == 0 {
			tmp = append(tmp, strconv.Itoa(i))
		}
		rtnValue = append(rtnValue, strings.Join(tmp, ""))
	}

	return rtnValue
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
			result := FizzBuzz(strings.Split(chopped, " "))
			fmt.Fprintln(os.Stdout, strings.Join(result, " "))

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
