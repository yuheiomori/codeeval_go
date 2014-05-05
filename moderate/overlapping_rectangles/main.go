package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func UpperInitial(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

type Point struct {
	x int
	y int
}

type Rect struct {
	p1     Point
	p2     Point
	left   int
	right  int
	top    int
	bottom int
}

func NewRect(p1, p2 Point) *Rect {
	r := new(Rect)
	r.p1 = p1
	r.p2 = p2
	r.left = int(math.Min(float64(p1.x), float64(p2.x)))
	r.right = int(math.Max(float64(p1.x), float64(p2.x)))
	r.bottom = int(math.Min(float64(p1.y), float64(p2.y)))
	r.top = int(math.Max(float64(p1.y), float64(p2.y)))
	return r
}

func solve(s string) string {
	input := strings.Split(s, ",")
	x1, _ := strconv.Atoi(input[0])
	y1, _ := strconv.Atoi(input[1])
	x2, _ := strconv.Atoi(input[2])
	y2, _ := strconv.Atoi(input[3])
	x3, _ := strconv.Atoi(input[4])
	y3, _ := strconv.Atoi(input[5])
	x4, _ := strconv.Atoi(input[6])
	y4, _ := strconv.Atoi(input[7])

	rect1 := NewRect(Point{x1, y1}, Point{x2, y2})
	rect2 := NewRect(Point{x3, y3}, Point{x4, y4})

	is_horizontal_overlap := true
	is_vertical_overlap := true
	if (rect1.left > rect2.right) || (rect1.right < rect2.left) {
		is_horizontal_overlap = false
	}

	if (rect1.top < rect2.bottom) || (rect1.bottom > rect2.top) {
		is_vertical_overlap = false
	}

	if is_horizontal_overlap && is_vertical_overlap {
		return "True"
	} else {
		return "False"
	}

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
