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

type QueryBoard struct {
	board [256][256]int
}

func (q *QueryBoard) setCol(col, val int) {
	for i := 0; i < 256; i++ {
		q.board[i][col] = val
	}
}

func (q *QueryBoard) setRow(row, val int) {
	for i := 0; i < 256; i++ {
		q.board[row][i] = val
	}
}

func (q *QueryBoard) queryCol(col int) (sum int) {
	for i := 0; i < 256; i++ {
		sum += q.board[i][col]
	}
	return
}

func (q *QueryBoard) queryRow(row int) (sum int) {
	for i := 0; i < 256; i++ {
		sum += q.board[row][i]
	}
	return
}

func main() {
	flag.Parse()

	read_file, _ := os.OpenFile(flag.Arg(0), os.O_RDONLY, 0600)
	defer read_file.Close()
	reader := bufio.NewReader(read_file)

	qb := new(QueryBoard)

	for {
		switch line, err := reader.ReadString('\n'); err {
		case nil:
			chopped := line[0 : len(line)-1]
			args := strings.Split(chopped, " ")
			switch args[0] {
			case "SetRow":
				row, _ := strconv.Atoi(args[1])
				val, _ := strconv.Atoi(args[2])
				qb.setRow(row, val)
			case "SetCol":
				col, _ := strconv.Atoi(args[1])
				val, _ := strconv.Atoi(args[2])
				qb.setCol(col, val)
			case "QueryRow":
				row, _ := strconv.Atoi(args[1])
				fmt.Fprintln(os.Stdout, qb.queryRow(row))
			case "QueryCol":
				col, _ := strconv.Atoi(args[1])
				fmt.Fprintln(os.Stdout, qb.queryCol(col))
			}

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
