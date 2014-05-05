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

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func NewTree(left *Tree, value int, right *Tree) *Tree {
	t := &Tree{left, value, right}
	return t
}

func NewTreeFromArray(values []int) *Tree {
	var t *Tree

	for i := 0; i < len(values); i++ {
		t = insert(t, values[i])
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func findLCA(t *Tree, p, q int) *Tree {
	if t == nil {
		return nil
	}
	if (t.Left != nil && t.Left.Value == p) ||
		(t.Left != nil && t.Left.Value == q) ||
		(t.Right != nil && t.Right.Value == p) ||
		(t.Right != nil && t.Right.Value == q) {
		return t
	} else {
		l := findLCA(t.Left, p, q)
		r := findLCA(t.Right, p, q)

		if l != nil && r != nil {
			return t
		} else if l != nil {
			return l
		} else {
			return r
		}
	}
	return nil
}

func main() {
	flag.Parse()

	read_file, _ := os.OpenFile(flag.Arg(0), os.O_RDONLY, 0600)
	defer read_file.Close()
	reader := bufio.NewReader(read_file)

	t := NewTreeFromArray([]int{30, 8, 52, 3, 20, 10, 29})

	for {
		switch line, err := reader.ReadString('\n'); err {
		case nil:
			chopped := line[0 : len(line)-1]
			integers := strings.Split(chopped, " ")
			p1, _ := strconv.Atoi(integers[0])
			p2, _ := strconv.Atoi(integers[1])
			fmt.Fprintln(os.Stdout, findLCA(t, p1, p2).Value)

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
