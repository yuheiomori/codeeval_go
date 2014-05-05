package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	read_file, _ := os.OpenFile(flag.Arg(0), os.O_RDONLY, 0600)
	defer read_file.Close()

	if fi, err := read_file.Stat(); err == nil {
		fmt.Fprintln(os.Stdout, fi.Size())
	}
}
