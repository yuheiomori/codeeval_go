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

func Intersection(l1, l2 []int) (result []int) {

	for i, j := 0, 0; i < len(l1) && j < len(l2); {
		if l1[i] < l2[j] {
			i++
		} else if l2[j] < l1[i] {
			j++
		} else {
			result = append(result, l2[j])
			j++
			i++
		}
	}
	return
}

func Convert(l []string) (result []int) {
	for i := 0; i < len(l); i++ {
		converted, _ := strconv.Atoi(l[i])
		result = append(result, converted)
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
			lists := strings.Split(chopped, ";")
			l1 := Convert(strings.Split(lists[0], ","))
			l2 := Convert(strings.Split(lists[1], ","))
			result := Intersection(l1, l2)
			result_s := make([]string, len(result))
			for i := 0; i < len(result); i++ {
				result_s[i] = strconv.Itoa(result[i])
			}
			fmt.Fprintln(os.Stdout, strings.Join(result_s, ","))
		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
