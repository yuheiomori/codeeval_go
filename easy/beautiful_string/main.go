package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
)

type runeCountMap struct {
	m map[rune]int
	s []rune
}

func (rcm *runeCountMap) Len() int {
	return len(rcm.m)
}

func (rcm *runeCountMap) Less(i, j int) bool {
	return rcm.m[rcm.s[i]] > rcm.m[rcm.s[j]]
}

func (rcm *runeCountMap) Swap(i, j int) {
	rcm.s[i], rcm.s[j] = rcm.s[j], rcm.s[i]
}

func sortedKeys(m map[rune]int) []rune {
	rcm := new(runeCountMap)
	rcm.m = m
	rcm.s = make([]rune, len(m))
	i := 0
	for key, _ := range m {
		rcm.s[i] = key
		i++
	}
	sort.Sort(rcm)
	return rcm.s
}

func getScore(s string) (score int) {
	memo := make(map[rune]int)
	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}
		memo[unicode.ToLower(r)]++
	}

	point := 26
	for _, k := range sortedKeys(memo) {
		score += point * memo[k]
		point--
		if point == 0 {
			break
		}
	}
	return score

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
			result := getScore(chopped)
			fmt.Fprintln(os.Stdout, result)

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
