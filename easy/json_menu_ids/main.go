package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type Item struct {
	Id    int
	Label *string
}

type Menu struct {
	Header *string
	Items  []*Item
}

type Message struct {
	Menu Menu
}

func solve(s string) (r int) {
	var m Message
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range m.Menu.Items {
		if item != nil && item.Label != nil {
			r += item.Id
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
			if chopped != "" {
				result := solve(chopped)
				fmt.Fprintln(os.Stdout, result)
			}

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
