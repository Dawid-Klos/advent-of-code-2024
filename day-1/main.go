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
	input, err := os.Open("input.txt")
	check(err)
	defer input.Close()

	r := bufio.NewReader(input)

	line, err := r.ReadString('\n')
	check(err)

	line, ok := strings.CutSuffix(line, "\n")
	if !ok {
		fmt.Println("Error removing suffix from the line:\n", line)
	}

	nums := strings.Split(line, "   ")
	num1, err := strconv.Atoi(nums[0])
	check(err)

	num2, err := strconv.Atoi(nums[1])
	check(err)

	fmt.Println(nums)
	fmt.Println("num1: ", num1)
	fmt.Println("num2: ", num2)
}
