package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInput() (string, error) {
	input :=""
	_, err := fmt.Fscan(os.Stdin, &input)
	if err != nil {
		return "", err
	}
	return input, nil
}

func makeArray(s string) ([]float64, error) {
	a := []float64{}
	ss := strings.Split(s, ",")
	for _, val := range ss {
		c, err := strconv.ParseFloat(val,64)
		if err != nil {
			return nil, err
		}
		a = append(a, c)
	}
	return a, nil
}

func convertStoA(s *[]float64) (res [3]float64) {
	for i, v := range *s {
		res[i] = v
	}
	return
}

func getTroika(slice *[]float64) (resMap map[[3]float64]int) {
	resMap = make(map[[3]float64]int)
	for i, val1 := range *slice {
		for j, val2 := range *slice {
			for k, val3 := range *slice {
				if val1+val2+val3 == 0 && i != j && i != k && j != k {
					arr := []float64{val1, val2, val3}
					sort.Float64s(arr)
					resArray := convertStoA(&arr)

					if _, ok := resMap[resArray]; !ok {
						resMap[resArray] = 1
					}
				}
			}
		}
	}
	return
}

func myOutput(amd map[[3]float64]int) {
	result := ""
	for k, _ := range amd {
		for i, val := range k {
			if i < 2 {
				result += fmt.Sprintf("%v",val) + ","
			} else {
				result += fmt.Sprintf("%v",val) + " "
			}

		}

	}
	result = strings.TrimSuffix(result, " ")
	fmt.Println(result)
}

func main() {

	input, errInput := getInput()
	if errInput != nil {

	}
	floats, err := makeArray(input)
	if err != nil {

	}
	amd := getTroika(&floats)
	myOutput(amd)

}
