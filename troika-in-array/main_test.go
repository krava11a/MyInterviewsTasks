package main

import (
	"testing"
)

func MyEqualSlices(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestMakeArrayNulls(t *testing.T) {
	input := "0,0,0,0"
	expected := []float64{0, 0, 0, 0}
	got, _ := makeArray(input)
	if !MyEqualSlices(got, expected) {

		t.Errorf("Input 0,0,0,0 , expected:  , result:%v", got)
	} else {
		t.Logf("Success!!! Input 0,0,0,0 , expected: [0 0 0 0] , result:%v\n", got)

	}

	//refactoring
}

func TestMakeArrayInts(t *testing.T) {
	input := "-1,0,1,2,-1,-4"
	expected := []float64{-1, 0, 1, 2, -1, -4}
	got, _ := makeArray(input)
	if !MyEqualSlices(got, expected) {
		t.Errorf("Input -1,0,1,2,-1,-4 , expected: [-1 0 1 2 -1 -4]  , result:%v", got)
	} else {
		t.Logf("Success!! Input -1,0,1,2,-1,-4 , expected: [-1 0 1 2 -1 -4] , result:%v\n", got)
	}
	//refactoring
}

func TestMakeArrayFloats(t *testing.T) {
	input := "-1.1,0.7,1.1,-2,1,2,-1,-4,1.2,2.2"
	expected := []float64{-1.1, 0.7, 1.1, -2, 1, 2, -1, -4, 1.2, 2.2}
	got, _ := makeArray(input)
	if !MyEqualSlices(got, expected) {
		t.Errorf("Input -1,0,1,2,-1,-4 , expected:  , result:%v", got)
	} else {
		t.Logf("Success!! Input -1.1, 0.7, 1.1, -2, 1, 2, -1, -4, 1.2, 2.2 , expected: [-1.1 0.7 1.1 -2 1 2 -1 -4 1.2 2.2] , result:%v\n", got)
	}
	//refactoring
}
