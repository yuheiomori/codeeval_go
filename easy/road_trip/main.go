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

func parseDistances(s string) []int {
	result := []int{0}
	points := strings.Split(s, ";")
	for _, point := range points {
		point_info := strings.Split(strings.Trim(point, " "), ",")
		if len(point_info) != 2 {
			continue
		}
		distance, _ := strconv.Atoi(point_info[1])
		result = append(result, distance)
	}
	sort.Ints(result)
	return result
}

func roadTrip(s string) (result string) {

	distances := parseDistances(s)

	tmp := make([]string, 0)

	for i := 0; i < len(distances)-1; i++ {
		tmp = append(tmp, strconv.Itoa(distances[i+1]-distances[i]))
	}

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
			result := roadTrip(chopped)
			fmt.Fprintln(os.Stdout, result)

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
