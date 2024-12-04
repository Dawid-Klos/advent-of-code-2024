package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var leftList []int
	var rightList []int

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		line, ok := strings.CutSuffix(line, "\n")
		if !ok {
			fmt.Println("Error removing suffix from the line:\n", line)
		}

		nums := strings.Split(line, "   ")
		num1, err := strconv.Atoi(nums[0])
		check(err)

		num2, err := strconv.Atoi(nums[1])
		check(err)
		leftList = append(leftList, num1)
		rightList = append(rightList, num2)
	}

	sort.Slice(leftList, func(i, j int) bool { return leftList[i] < leftList[j] })
	sort.Slice(rightList, func(i, j int) bool { return rightList[i] < rightList[j] })

	var result int

	for i, num := range leftList {
		if num > rightList[i] {
			result += num - rightList[i]
		} else {
			result += rightList[i] - num
		}
	}

	fmt.Println("result: ", result)

}
