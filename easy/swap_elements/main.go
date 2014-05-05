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

func solve(s string) (r string) {
	input := strings.Split(s, ":")
	swap_target_array := strings.Split(strings.Trim(input[0], " "), " ")
	swap_positions := strings.Split(strings.Trim(input[1], " "), ",")

	for _, swap_position := range swap_positions {
		swap_pos := strings.Split(strings.Trim(swap_position, " "), "-")
		swap_pos1, _ := strconv.Atoi(swap_pos[0])
		swap_pos2, _ := strconv.Atoi(swap_pos[1])
		swap_target_array[swap_pos1], swap_target_array[swap_pos2] = swap_target_array[swap_pos2], swap_target_array[swap_pos1]
	}
	return strings.Join(swap_target_array, " ")
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
			//			fmt.Fprintln(os.Stdout, chopped)
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
