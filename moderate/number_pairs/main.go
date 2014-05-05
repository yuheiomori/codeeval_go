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

type StringArray []string

func (sa *StringArray) IntArray() []int {
	result := make([]int, len(*sa))
	for idx, c := range *sa {
		i, _ := strconv.Atoi(c)
		result[idx] = i
	}
	return result
}

type Pair [2]int

func (p *Pair) First() int {
	return p[0]
}

func (p *Pair) Second() int {
	return p[1]
}

func (p *Pair) String() (s string) {
	return strings.Join([]string{
		strconv.Itoa(p.First()),
		strconv.Itoa(p.Second())},
		",")
}

type Pairs []Pair

func (ps *Pairs) String() string {

	if len(*ps) == 0 {
		return "NULL"
	}
	tmp := make([]string, len(*ps))

	for i, p := range *ps {
		tmp[i] = p.String()
	}
	return strings.Join(tmp, ";")

}
func number_pairs(l []int, x int) Pairs {

	memo := make(map[int]bool)

	var inner_func func(ll []int, xx int) (rr Pairs)
	inner_func = func(ll []int, xx int) (rr Pairs) {
		top, rest := ll[0], ll[1:]

		if !memo[top] {
			for _, d := range rest {
				if top+d == xx {
					rr = append(rr, Pair{top, d})
					memo[d] = true
					break
				}
			}
		}

		if len(rest) > 1 {
			return append(rr, inner_func(rest, x)...)
		}
		return rr

	}
	return inner_func(l, x)

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
			var splitted StringArray = strings.Split(chopped, ";")
			var stringArray StringArray = strings.Split(splitted[0], ",")
			x, _ := strconv.Atoi(splitted[1])
			pairs := number_pairs(stringArray.IntArray(), x)
			fmt.Fprintln(os.Stdout, pairs.String())

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
