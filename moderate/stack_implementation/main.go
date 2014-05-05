package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type Element struct {
	value interface{}
	next  *Element
}

type Stack struct {
	top  *Element
	size int
	lock *sync.Mutex
}

func NewStack() *Stack {
	stack := &Stack{}
	stack.lock = &sync.Mutex{}
	return stack
}

func (stack *Stack) Len() int {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	return stack.size
}

func (stack *Stack) Push(value interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	stack.top = &Element{value, stack.top}
	stack.size++
}

func (stack *Stack) Pop() (value interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size > 0 {
		value, stack.top = stack.top.value, stack.top.next
		stack.size--
		return
	}
	return nil
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
			chars := strings.Split(chopped, " ")

			stack := NewStack()
			for _, c := range chars {
				stack.Push(c)
			}

			alt_flg := true
			result := make([]string, 0)
			for stack.Len() > 0 {
				if alt_flg {
					result = append(result, stack.Pop().(string)) // cast to string
				} else {
					stack.Pop() // throw away
				}
				alt_flg = !alt_flg
			}
			fmt.Fprintln(os.Stdout, strings.Join(result, " "))

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
