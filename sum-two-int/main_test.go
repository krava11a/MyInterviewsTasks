package main

import "testing"

func TestRightDigits(t *testing.T) {
	sum, _ := sumInts("2", "3")
	if sum != "5" {
		t.Errorf("Input 2 and 3 , expected: 5 , result:%v", sum)
	}
}

func TestString(t *testing.T) {
	sum, _ := sumInts("11111", "444@")
	if sum != "not digit" {
		t.Errorf("Input 11111 and 444@ , expected: not digit , result:%v", sum)
	}
}

func TestString2(t *testing.T) {
	sum, _ := sumInts("sadf", "sdas@")
	if sum != "not digit" {
		t.Errorf("Input sadf and sdas@ , expected: not digit , result:%v", sum)
	}
}

func TestNull(t *testing.T) {
	sum, _ := sumInts("0", "0")
	if sum != "not digit" {
		t.Errorf("Input 0 and 0 , expected: not digit , result:%v", sum)
	}
}

func TestNullInt(t *testing.T) {
	sum, _ := sumInts("$", "0")
	if sum != "not digit" {
		t.Errorf("Input $ and 0 , expected: not digit , result:%v", sum)
	}
}

func TestBigDigit(t *testing.T) {
	sum, _ := sumInts("231234123412341234123412", "1")
	if sum != "9223372036854775808" {
		t.Errorf("Input 231234123412341234123412 and 1 , expected: 231234123412341234123413 , result:%v", sum)
	}
}

func TestBigDigit2(t *testing.T) {
	sum, _ := sumInts("1","231234123412341234123412")
	if sum != "9223372036854775808" {
		t.Errorf("Input 231234123412341234123412 and 1 , expected: 231234123412341234123413 , result:%v", sum)
	}
}

func TestMinus(t *testing.T) {
	sum, _ := sumInts("-23412", "1")
	if sum != "not digit" {
		t.Errorf("Input -23412 and 1 , expected: not digit , result:%v", sum)
	}
}
