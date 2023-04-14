package main

import "fmt"

func main() {
	fmt.Println(Add(2,3))
	fmt.Println(Sum([]int{1, 2, 3, 4, 5}))
	fmt.Println(Average([]int{1, 2, 3, 4, 5}))
}

func Add(a int, b int) int {
	return a + b
}

func Sum(numbers []int) int {
	res := 0

	for _, num := range numbers {
		res += num
	}

	return res
}

func Average(numbers []int) int {
	res := 0

	for _, num := range numbers {
		res += num
	}

	return res / len(numbers)
}