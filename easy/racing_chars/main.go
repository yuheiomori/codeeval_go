package main

import (
	"fmt"
	"log"
	"strings"
)
import "bufio"
import "os"

const (
	PASSING_TYPE_STRAIGHT   = "|"
	PASSING_TYPE_LEFT_TURN  = "\\"
	PASSING_TYPE_RIGHT_TURN = "/"
	ROAD_TYPE_GATE          = "_"
	ROAD_TYPE_CHECKPOINT    = "C"
)

func findNext(line string) (int, string) {
	gate_pos := strings.Index(line, ROAD_TYPE_GATE)
	checkpoint_pos := strings.Index(line, ROAD_TYPE_CHECKPOINT)

	var next_pos int
	var next_type string

	if checkpoint_pos != -1 {
		next_pos = checkpoint_pos
		next_type = ROAD_TYPE_CHECKPOINT
	} else {
		next_pos = gate_pos
		next_type = ROAD_TYPE_GATE
	}
	return next_pos, next_type

}

func decideWayOfPassing(last_pos, next_pos int) string {
	var way_of_passing string
	way_of_passing = PASSING_TYPE_STRAIGHT

	if last_pos != -1 {
		if next_pos > last_pos {
			way_of_passing = PASSING_TYPE_LEFT_TURN
		} else if next_pos < last_pos {
			way_of_passing = PASSING_TYPE_RIGHT_TURN
		}
	}
	return way_of_passing
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var last_pos int = -1

	for scanner.Scan() {
		line := scanner.Text()

		next_pos, next_type := findNext(line)
		way_of_passing := decideWayOfPassing(last_pos, next_pos)
		last_pos = next_pos

		fmt.Println(strings.Replace(line, next_type, way_of_passing, 1))

	}
}
