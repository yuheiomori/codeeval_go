package main

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)


func TestStringArray(t *testing.T) {
	Convey("Convert To IntArray", t, func() {
		array1 := StringArray{"1", "2", "3"}
		So(array1.IntArray(), ShouldResemble, []int{1, 2, 3})
	})
}
func TestPair(t *testing.T) {
	Convey("String representation", t, func() {
		pair := Pair{2, 5}
		So(pair.String(), ShouldEqual, "2,5")
	})
}
func TestPairs(t *testing.T) {
	Convey("String representation for empty", t, func() {
		pairs := Pairs{}
		So(pairs.String(), ShouldEqual, "NULL")
	})
	Convey("String representation for single pair", t, func() {
		pairs := Pairs{Pair{1, 2}}
		So(pairs.String(), ShouldEqual, "1,2")

	})
	Convey("String representation for some pairs", t, func() {
		pairs := Pairs{Pair{1, 2}, Pair{3, 4}}
		So(pairs.String(), ShouldEqual, "1,2;3,4")
	})
}
func TestNumberPairs(t *testing.T) {
	Convey("Case1", t, func() {
		result := number_pairs([]int{1, 2, 3, 4, 6}, 5)
		So(result, ShouldResemble, Pairs{Pair{1, 4}, Pair{2, 3}})
		So(result.String(), ShouldEqual, "1,4;2,3")

	})
	Convey("Case2", t, func() {
		result := number_pairs([]int{2, 4, 5, 6, 9, 11, 15}, 20)
		So(result, ShouldResemble, Pairs{Pair{5, 15}, Pair{9, 11}})
		So(result.String(), ShouldEqual, "5,15;9,11")
	})
	Convey("Case3", t, func() {
		result := number_pairs([]int{1, 2, 3, 4}, 50)
		cValue := reflect.ValueOf(result)
		fmt.Println(cValue)
		aValue := reflect.ValueOf(Pairs{})
		fmt.Println(aValue)
		So(result, ShouldEqual, Pairs{})
		So(result.String(), ShouldEqual, "NULL")
	})
}

