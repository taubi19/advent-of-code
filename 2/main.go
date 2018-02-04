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

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		// find the difference between the smallest and biggest number in words
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
		sum += (max - min)
	}
	fmt.Println(sum)
}
