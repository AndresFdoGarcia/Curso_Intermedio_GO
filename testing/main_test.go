package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	// sum := sum(1, 2)
	// if sum != 3 {
	// 	t.Errorf("Sum was incorrect, got: %d, expected: %d.", sum, 3)
	// }
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
	}

	for _, item := range tables {
		total := sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrect, got: %d, expected: %d.", total, item.n)
		}
	}
}

func TestGetMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 2},
		{2, 1, 2},
		{3, 3, 3},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.n {
			t.Errorf("GetMax was incorrect, got: %d, expected: %d.", max, item.n)
		}
	}
}

func TestFibonacci(t *testing.T) {
	table := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{3, 2},
		{50, 12586269025},
	}

	for _, item := range table {
		fib := Fibonacci(item.a)
		if fib != item.n {
			t.Errorf("Fibonacci was incorrect, got: %d, expected: %d.", fib, item.n)
		}
	}
}
