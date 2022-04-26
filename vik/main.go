package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumInts(a, b string) string {
	xa, err := strconv.Atoi(a)
	if err != nil || xa <= 0 {
		return "not digit"
	}
	xb, err := strconv.Atoi(b)
	if err != nil || xb <= 0 {
		return "not digit"
	}
	sum := strconv.Itoa(xa + xb)
	return sum
}

func getInput() (string, error) {
	var reader = bufio.NewReader(os.Stdin)
	name, error := reader.ReadString('\n')
	return name, error
}

func splitInput(input string) (string, string, error) {

	rez := strings.Split(input, " ")
	if len(rez) > 1 {

		return rez[0], strings.Replace(rez[1],"\n","",1), nil
	} else {
		return "", "", errors.New("not digit")
	}
}
func main() {

	input, errInput := getInput()
	if errInput == nil {
		val1, val2, errParse := splitInput(input)
		if errParse == nil {
			sum := sumInts(val1, val2)
			res, err := strconv.Atoi(sum)
			if err != nil {
				fmt.Println("not digit")
			}else {
				fmt.Println(res)
			}

		} else {
			fmt.Println(errParse)
		}

	} else {
		fmt.Println("not digit")
	}

}