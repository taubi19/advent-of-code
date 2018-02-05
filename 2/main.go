package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	check(err)
	scanner := bufio.NewScanner(file)

	sum := 0
	sum2 := 0
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		// find the difference between the smallest and biggest number in words
		sum += diffMaxMin(words)
		sum2 += divisibleValRes(words)
	}
	fmt.Printf("Sum1 = %v\n", sum)
	fmt.Printf("Sum2 = %v\n", sum2)

}

// Accepts an array of strings that are supposed to be ints. Each element is then checked if it has a even divisor
// in the array. When such a divisor is found the division result is added to linesum. When all elements have been
// checked the sum is returned.
func divisibleValRes(words []string) int {
	var linesum = 0
	for i := 0; i < len(words); i++ {
		eli, err := strconv.Atoi(words[i])
		check(err)
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			elj, err := strconv.Atoi(words[j])
			check(err)
			// if el is equally divided by words[i]
			if eli%elj == 0 {
				fmt.Printf("eli = %v, elj = %v, eli/elj= %v, eli mod elj= %v \n", eli, elj, eli/elj, eli%elj)
				linesum += eli / elj
				continue
			}
			// the other way around, words[i] is equally divided by el
			if elj%eli == 0 {
				fmt.Printf("other way: eli = %v, elj = %v, eli/elj= %v, eli mod elj= %v \n", eli, elj, elj/eli, elj%eli)
				linesum += eli / elj
				continue
			}
		}
	}
	return linesum
}

// Accepts an array of strings that are expected to be ints. We search the array for the minimum and maximum, then
// the difference between those two is calculated and returned.
func diffMaxMin(words []string) int {

	firstEl, err := strconv.Atoi(words[0])
	check(err)
	var min, max int = firstEl, firstEl
	for i := 0; i < len(words); i++ {
		num, err := strconv.Atoi(words[i])
		check(err)
		if min > num {
			min = num
		}
		if max < num {
			max = num
		}
	}
	return (max - min)
}
