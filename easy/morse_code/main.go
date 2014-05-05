package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	morse2alpha = map[string]string{
		".-":    "A",
		"-.-.":  "C",
		"-...":  "B",
		".":     "E",
		"-..":   "D",
		"--.":   "G",
		"..-.":  "F",
		"..":    "I",
		"....":  "H",
		"-.-":   "K",
		".---":  "J",
		"--":    "M",
		".-..":  "L",
		"---":   "O",
		"-.":    "N",
		"--.-":  "Q",
		".--.":  "P",
		"...":   "S",
		".-.":   "R",
		"-":     "T",
		"..-":   "U",
		".--":   "W",
		"...-":  "V",
		"-.--":  "Y",
		"-..-":  "X",
		"--..":  "Z",
		".----": "1",
		"-----": "0",
		"...--": "3",
		"..---": "2",
		".....": "5",
		"....-": "4",
		"--...": "7",
		"-....": "6",
		"----.": "9",
		"---..": "8",
	}
)

func morseCode(s string) (result string) {
	for _, e := range strings.Split(s, " ") {
		if e == "" {
			result += " "
		}
		result += morse2alpha[e]
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
			result := morseCode(chopped)
			fmt.Fprintln(os.Stdout, result)

		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)

		}
	}
}
