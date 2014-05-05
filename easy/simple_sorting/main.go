package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func convert(in []string) (out []float64) {
	for _, s := range in {
		f, _ := strconv.ParseFloat(s, 32)
		out = append(out, f)
	}
	return
}

func printResult(l []float64) {
	tmp := make([]string, len(l))
	for idx, f := range l {
		tmp[idx] = strconv.FormatFloat(f, 'f', 3, 64)
	}
	fmt.Fprintln(os.Stdout, strings.Join(tmp, " "))
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
			splitted := strings.Split(chopped, " ")
			float_list := convert(splitted)
			sort.Float64s(float_list)
			printResult(float_list)

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
