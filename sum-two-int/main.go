package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)


func sumInts(a, b string) (string, error) {
	xa, err := strconv.ParseUint(a, 10, 64)
	if err != nil || xa <= 0 {
		if numErr, ok := err.(*strconv.NumError); ok {
			if numErr.Err == strconv.ErrRange {
				inf := strconv.FormatUint(uint64(math.Inf(1)), 10)
				return inf, err
			}
			if numErr.Err == strconv.ErrSyntax {
				return "not digit", nil
			}
		} else {
			return "not digit", nil
		}

	}

	xb, err := strconv.ParseUint(b, 10, 64)
	if err != nil || xb <= 0 {
		if numErr, ok := err.(*strconv.NumError); ok {
			if numErr.Err == strconv.ErrRange {
				inf := strconv.FormatUint(uint64(math.Inf(1)), 10)
				return inf, err
			}
			if numErr.Err == strconv.ErrSyntax {
				return "not digit", nil
			}
		} else {
			return "not digit", nil
		}
	}
	sum := strconv.FormatUint(xa+xb, 10)
	return sum, nil
}

func main() {
	input1 := ""
	input2 := ""

	_, err := fmt.Fscan(os.Stdin, &input1)
	if err != nil {
		fmt.Println("not digit")

	}
	_, err = fmt.Fscan(os.Stdin, &input2)
	if err != nil {
		fmt.Println("not digit")
	}
	sum, err := sumInts(input1, input2)
	if err!=nil {
		fmt.Println(math.Inf(1))
	}else {
		fmt.Println(sum)
	}


}
