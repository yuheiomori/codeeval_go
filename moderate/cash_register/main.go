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

var COINS map[string]float64 = map[string]float64{
	"PENNY":       .01,
	"NICKEL":      .05,
	"DIME":        .10,
	"QUARTER":     .25,
	"HALF DOLLAR": .50,
	"ONE":         1.00,
	"TWO":         2.00,
	"FIVE":        5.00,
	"TEN":         10.00,
	"TWENTY":      20.00,
	"FIFTY":       50.00,
	"ONE HUNDRED": 100.00,
}

type sortedMap struct {
	m map[string]float64
	s []string
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}

func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func sortedKeys(m map[string]float64) []string {
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}

func solve(s string) (result []string) {
	prices := strings.Split(s, ";")
	purchase_price, _ := strconv.ParseFloat(prices[0], 64)
	cash, _ := strconv.ParseFloat(prices[1], 64)

	if purchase_price == cash {
		result = append(result, "ZERO")
	} else if purchase_price > cash {
		result = append(result, "ERROR")
	} else {
		change := int(cash*100) - int(purchase_price*100)
		for _, coin_name := range sortedKeys(COINS) {
			coin_value := int(COINS[coin_name] * 100)

			if change >= coin_value {
				num := change / coin_value
				change -= coin_value * num
				for i := 0; i < num; i++ {
					result = append(result, coin_name)
				}
			}
			if change == 0 {
				break
			}
		}
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
			if len(chopped) > 0 {
				result := solve(chopped)
				fmt.Fprintln(os.Stdout, strings.Join(result, ","))
			}

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
