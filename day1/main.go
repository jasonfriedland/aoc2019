/*
https://adventofcode.com/2019/day/1
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calculate(in int64) int64 {
	// Divide by three, round down, and subtract 2
	return int64(float64(in/3) - 2)
}

// reduce recursively calls calculate on the input.
func reduce(x int64, y int64) (int64, int64) {
	x = calculate(x)
	if x < 1.0 {
		return 0.0, y
	}
	y += x
	return reduce(x, y)
}

func main() {
	f, _ := os.Open("day1/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var modTotal, fuelTotal int64
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		// Part 1
		modTotal += calculate(int64(i))
		// Part 2
		_, fuel := reduce(int64(i), 0)
		fuelTotal += fuel
	}
	// Answers
	fmt.Println(modTotal, fuelTotal)
}
