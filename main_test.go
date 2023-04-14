package main

import "testing"

func TestAdd(t *testing.T) {
	var expected int = 5
	var actual int = Add(2,3)

	if expected != actual {
		t.Errorf("Expected: %v got: %v", expected, actual)
	}
}

func TestSum(t *testing.T) {
    expected := 15
    actual := Sum([]int{1, 2, 3, 4, 5})

    if actual != expected {
        t.Errorf("Expected: %v got: %v", expected, actual)
    }
}

func TestAverage(t *testing.T) {
    expected := 3
    actual := Average([]int{1, 2, 3, 4, 5})

    if actual != expected {
        t.Errorf("Expected: %v got: %v", expected, actual)
    }
}