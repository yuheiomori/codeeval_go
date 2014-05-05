package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// リストからidx番目の要素を抜かしたものを取得
func delete(idx int, L []rune) (result []rune) {
	result = append(result, L[:idx]...)
	result = append(result, L[idx+1:]...)
	return
}

// sliceの全組み合わせを返す
func Permute(L []rune) (result [][]rune) {
	var inner func(LL []rune)
	inner = func(LL []rune) {
		if len(LL) == 0 {
			result = append(result, []rune{}) //
		}
		for idx, r := range L {
			for _, t := range Permute(delete(idx, L)) {
				result = append(result, append([]rune{r}, t...))
			}
		}
	}
	inner(L)
	return

}

func solve(s string) string {
	tmp := sort.StringSlice{}
	for _, l := range Permute([]rune(s)) {
		tmp = append(tmp, string(l))
	}
	tmp.Sort()
	return strings.Join(tmp, ",")
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
