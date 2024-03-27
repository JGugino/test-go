package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	PARSE_BASE = 10
	BIT_SIZE   = 64
)

func main() {
	values := []float64{
		264.26,
		156486,
		2156.54,
		418543484,
		13648.45,
		16,
	}

	for i := 0; i < len(values); i++ {
		testCurrencyConversion(values[i])
	}
}

func testCurrencyConversion(value float64) bool {
	var converted int64

	converted, _ = ConvertUSDDollarToPennies(value)

	if (converted / 100) == int64(value) {
		fmt.Printf("Conversion Correct: %d(converted) == %.2f(original) \n", converted, value)
		return true
	}

	fmt.Printf("Conversion Incorrect: %d(converted) != %.2f(original) \n", converted, value)
	return false
}

func ConvertUSDDollarToPennies(dollarAmount float64) (int64, error) {

	stringAmount := fmt.Sprintf("%.2f", dollarAmount)
	splitString := strings.Split(stringAmount, ".")

	secondNum, err := strconv.ParseInt(splitString[1], PARSE_BASE, BIT_SIZE)

	if err != nil {
		return -1, err
	}

	if secondNum <= 0 {
		splitString[0] += "00"

		num, err := strconv.ParseInt(splitString[0], PARSE_BASE, BIT_SIZE)

		if err != nil {
			return -1, err
		}

		return num, nil
	}

	removedDecimal := strings.Replace(stringAmount, ".", "", -1)

	num, err := strconv.ParseInt(removedDecimal, PARSE_BASE, BIT_SIZE)

	if err != nil {
		return -1, err
	}

	return num, nil
}
