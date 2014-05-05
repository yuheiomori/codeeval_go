package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func isArmstringNumber(i, n int) bool {

	sum := 0
	temp := i

	for {
		remainder := temp % 10
		sum = sum + int(math.Pow(float64(remainder), float64(n)))
		temp = temp / 10
		if temp == 0 {
			break
		}
	}

	return i == sum
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
			n := len(chopped)
			i, _ := strconv.Atoi(chopped)
			result := isArmstringNumber(i, n)
			if result {
				fmt.Fprintln(os.Stdout, "True")
			} else {
				fmt.Fprintln(os.Stdout, "False")
			}

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
