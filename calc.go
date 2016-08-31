package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type mathFn func(float64, float64) float64

var topLevelOperators = map[string]mathFn{
	"+": add,
	"-": subtract,
}

var lowLevelOperators = map[string]mathFn{
	"*": multiply,
	"/": divide,
}

func main() {
	args := os.Args[1:]
	fmt.Println(reduce(strings.Split(args[0], " ")))
}

func reduce(rest []string) float64 {
	operatorIndex, hasOperator := indexOfFirstTopLevelOperator(rest)
	if hasOperator == false {
		return combineLowLevelOperators(rest)
	} else {
		next := operatorIndex + 1
		method := topLevelOperators[rest[operatorIndex]]
		left := reduce(rest[:operatorIndex])
		right := reduce(rest[next:])
		return method(left, right)
	}
}

func combineLowLevelOperators(rest []string) float64 {
	if len(rest) == 1 {
		x, _ := strconv.Atoi(rest[0])
		return float64(x)
	}
	method := lowLevelOperators[rest[1]]
	left := combineLowLevelOperators(rest[:1])
	right := combineLowLevelOperators(rest[2:])
	return method(left, right)
}

func indexOfFirstTopLevelOperator(rest []string) (int, bool) {
	for i := len(rest) - 1; i >= 0; i-- {
		v := rest[i]
		_, hasKey := topLevelOperators[v]
		if hasKey == true {
			return i, true
		}
	}
	return 0, false
}

func add(a float64, b float64) float64 {
	return a + b
}

func subtract(a float64, b float64) float64 {
	return a - b
}

func multiply(a float64, b float64) float64 {
	return a * b
}

func divide(a float64, b float64) float64 {
	return float64(a) / float64(b)
}
