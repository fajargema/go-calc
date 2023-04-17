package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	fmt.Println(Add(2,3))
	fmt.Println(Sum([]int{1, 2, 3, 4}))
	fmt.Println(Average([]int{1, 2, 3, 4, 5}))
	e.Logger.Fatal(e.Start(":8080"))
	
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