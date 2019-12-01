package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	// specs
	/*
		- divide by three
		- round down
		- substract by two
		- sum the result for each value
	*/
	f, err := os.Open("aoc.txt")
	if err != nil {
		panic(err)
	}
	numbers := reader(f)
	var sumit float64
	// var sumValues float64
	var r float64
	var tmp []float64

	for _, mass := range numbers {
		sumit += calculateFuel(mass)
		r += sum(reduceFuel(mass, tmp))
	}
	fmt.Printf("The answer is: %f - the second quiz is: %f", sumit, r)
}

func calculateFuel(mass float64) float64 {
	res := math.Floor(mass/3) - 2
	if res > 0 {
		return res
	}
	return 0
}

func reader(file io.Reader) []float64 {
	var numbers []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, value)
	}
	return numbers
}
