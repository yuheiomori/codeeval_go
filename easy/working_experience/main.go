package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
	"time"
)

type TimeInfo struct {
	Epoch int64
	Start bool
}

type TimeInfoList []TimeInfo

func (tl TimeInfoList) Len() int {
	return len(tl)
}

func (tl TimeInfoList) Swap(i, j int) {
	tl[i], tl[j] = tl[j], tl[i]
}

func (tl TimeInfoList) Less(i, j int) bool {
	return tl[i].Epoch < tl[j].Epoch
}

func (tl TimeInfoList) CalcTotalMonths() int64 {
	sort.Sort(tl)

	var all_months float64
	stack := make([]int64, 0)

	for _, t := range tl {
		if t.Start {
			stack = append(stack, t.Epoch)
		} else {
			end_time := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if len(stack) == 0 {
				months_diff := math.Floor(((float64(t.Epoch-end_time) / 3.15569e7 * 12) + 0.5) + 1)
				all_months = all_months + months_diff
			}
		}
	}
	return int64(all_months)

}

func (tl TimeInfoList) CalcTotalYears() int {
	return int(tl.CalcTotalMonths() / 12)
}

func isASCIISpace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r'
}

func TrimString(s string) string {
	for len(s) > 0 && isASCIISpace(s[0]) {
		s = s[1:]
	}
	for len(s) > 0 && isASCIISpace(s[len(s)-1]) {
		s = s[:len(s)-1]
	}
	return s
}

func Solve(spans []string) int {
	time_list := make(TimeInfoList, 0)
	for _, span := range spans {
		times := strings.Split(span, "-")
		start, _ := time.Parse("Jan 2006", TrimString(times[0]))
		end, _ := time.Parse("Jan 2006", TrimString(times[1]))
		time_list = append(time_list, TimeInfo{start.Unix(), true}, TimeInfo{end.Unix(), false})
	}

	return time_list.CalcTotalYears()

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
		spans := strings.Split(line, ";")
		fmt.Println(Solve(spans))
	}
}
