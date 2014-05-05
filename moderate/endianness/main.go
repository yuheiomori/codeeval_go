package main

import (
	"fmt"
	"os"
	"unsafe"
)

func main() {
	var x uint16 = 0x0102
	switch *(*byte)(unsafe.Pointer(&x)) {
	case 0x01:
		fmt.Fprintln(os.Stdout, "BigEndian")
		os.Exit(0)
	case 0x02:
		fmt.Fprintln(os.Stdout, "LittleEndian")
		os.Exit(0)
	}
}
