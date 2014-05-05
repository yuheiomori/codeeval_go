package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	x int
	y int
}

func (p *Point) distance(other Point) int64 {
	return int64(math.Sqrt(math.Pow(float64(p.x-other.x), 2.0) + math.Pow(float64(p.y-other.y), 2.0)))
}

func solve(s string) (r int64) {

	re := regexp.MustCompile("[-?\\d]+")
	line := re.FindAllString(s, -1)

	x1, _ := strconv.Atoi(line[0])
	y1, _ := strconv.Atoi(line[1])
	x2, _ := strconv.Atoi(line[2])
	y2, _ := strconv.Atoi(line[3])

	point1 := Point{x1, y1}
	point2 := Point{x2, y2}

	return point1.distance(point2)
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
