package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	r := parseDistances("Rkbs,5453; Wdqiz,1245; Rwds,3890; Ujma,5589; Tbzmo,1303;")

	if len(r) != 5 {
		t.Error()
	}
	if r[0] != 5453 {
		t.Error()
	}
	if r[1] != 1245 {
		t.Error()
	}
	if r[2] != 3890 {
		t.Error()
	}
	if r[3] != 5589 {
		t.Error()
	}
	if r[4] != 1303 {
		t.Error()
	}

}
