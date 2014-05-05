package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Data struct {
	words []string
	hints []int
}

func (d Data) Len() int {
	return len(d.words)
}

func (d Data) Swap(i, j int) {
	d.words[i], d.words[j] = d.words[j], d.words[i]
	d.hints[i], d.hints[j] = d.hints[j], d.hints[i]
}

func (d Data) Less(i, j int) bool {
	return d.hints[i] < d.hints[j]
}

func convert(sl []string) []int {
	il := make([]int, len(sl))
	for i, s := range sl {
		il[i], _ = strconv.Atoi(s)
	}
	return il
}

func FindLost(il []int) int {
	copied := make([]int, len(il))
	copy(copied, il)

	sort.Ints(copied)
	for i, e := range copied {
		if e != i+1 {
			return i + 1
		}
	}
	return len(il) + 1

}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		splitted := strings.Split(line, ";")
		words := strings.Split(splitted[0], " ")
		hints := convert(strings.Split(splitted[1], " "))

		hints = append(hints, FindLost(hints))

		data := Data{words, hints}
		sort.Sort(data)

		fmt.Println(strings.Join(data.words, " "))
	}
}
