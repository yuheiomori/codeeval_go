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

func douleSquares(n int) (count int) {
	s := int(math.Sqrt(float64(n)))
	for i := 0; i < s+1; i++ {
		D := math.Sqrt(float64(n - i*i))
		if D >= float64(i) && D == math.Ceil(D) {
			count++
		}
	}
	return

}

func main() {
	flag.Parse()

	read_file, _ := os.OpenFile(flag.Arg(0), os.O_RDONLY, 0600)
	defer read_file.Close()
	reader := bufio.NewReader(read_file)

	reader.ReadString('\n')

	for {
		switch line, err := reader.ReadString('\n'); err {
		case nil:
			chopped := line[0 : len(line)-1]
			n, _ := strconv.Atoi(chopped)
			result := douleSquares(n)
			fmt.Fprintln(os.Stdout, result)

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
