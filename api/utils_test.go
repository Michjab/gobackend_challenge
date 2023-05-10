package main

import (
	"math"
	"testing"
)

func TestIsOverflowMultiply1(t *testing.T) {
	x1 := 10
	x2 := 2
	x1x2 := x1 * x2
	expected := false

	gotResult := isOverflowMultiply(x1x2, x1, x2)

	if expected != gotResult {
		t.Errorf("Failed to check overflow multiplication for %v and %v. Expected value was %v, but got %v", x1, x2, expected, gotResult)
	}
}
func TestIsOverflowMultiply2(t *testing.T) {
	x1 := 10
	x2 := math.MaxInt
	x1x2 := x1 * x2
	expected := true

	gotResult := isOverflowMultiply(x1x2, x1, x2)

	if expected != gotResult {
		t.Errorf("Failed to check overflow multiplication for %v and %v. Expected value was %v, but got %v", x1, x2, expected, gotResult)
	}
}

func TestIsOverflowAdd1(t *testing.T) {
	x1 := 10
	x2 := math.MaxInt
	expected := true

	gotResult := isOverflowAdd(x1, x2)

	if expected != gotResult {
		t.Errorf("Failed to check overflow adding for %v and %v. Expected value was %v, but got %v", x1, x2, expected, gotResult)
	}
}

func TestIsOverflowAdd2(t *testing.T) {
	x1 := 10
	x2 := 1000
	expected := false

	gotResult := isOverflowAdd(x1, x2)

	if expected != gotResult {
		t.Errorf("Failed to check overflow adding for %v and %v. Expected value was %v, but got %v", x1, x2, expected, gotResult)
	}
}
func TestIsOverflowAdd3(t *testing.T) {
	x1 := 10
	x2 := math.MinInt
	expected := true

	gotResult := isOverflowAdd(x1, x2)

	if expected != gotResult {
		t.Errorf("Failed to check overflow adding for %v and %v. Expected value was %v, but got %v", x1, x2, expected, gotResult)
	}
}

func TestSumRow1(t *testing.T) {
	recordRow := []string{"1", "2", "3"}

	expected := 6

	gotResult, err := sumRow(recordRow)

	if expected != gotResult || err != nil {
		t.Errorf("Failed to check summing row for %v slice. Expected value was %v, but got %v with error %s", recordRow, expected, gotResult, err)
	}
}

func TestSumRow2(t *testing.T) {
	recordRow := []string{"notInt", "2", "3"}

	expected := 0

	gotResult, err := sumRow(recordRow)

	if expected != gotResult && err != nil {
		t.Errorf("Failed to check summing row for %v slice. Expected value was %v, but got %v with error %s", recordRow, expected, gotResult, err)
	}
}

func TestStrToInt1(t *testing.T) {
	recordRow := []string{"1", "2", "3"}

	expected := []int{1, 2, 3}

	gotResult, err := strToInt(recordRow)

	for i := range expected {
		if expected[i] != gotResult[i] || err != nil {
			t.Errorf("Failed to convert string to int row for %v slice. Expected slice was %v, but got %v with error %s", recordRow, expected, gotResult, err)
		}
	}
}

func TestMultiply1(t *testing.T) {
	recordRow := []int{4, 5, 6}

	expected := 120

	gotResult, err := multiply(recordRow)

	if expected != gotResult || err != nil {
		t.Errorf("Failed to convert string to int row for %v slice. Expected slice was %v, but got %v with error %s", recordRow, expected, gotResult, err)
	}
}

func TestMultiply2(t *testing.T) {
	recordRow := []int{0, 5, 6}

	expected := 0

	gotResult, err := multiply(recordRow)

	if expected != gotResult || err != nil {
		t.Errorf("Failed to convert string to int row for %v slice. Expected slice was %v, but got %v with error %s", recordRow, expected, gotResult, err)
	}
}

func TestTranspose(t *testing.T) {
	recordRow := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	expected := [][]string{
		{"1", "4", "7"},
		{"2", "5", "8"},
		{"3", "6", "9"},
	}

	gotResult, err := transpose(recordRow)

	for i := range expected {
		for j := range expected[i] {
			if expected[i][j] != gotResult[i][j] || err != nil {
				t.Errorf("Failed to convert string to int row for %v slice. Expected slice was %v, but got %v with error %s", recordRow, expected, gotResult, err)
			}
		}
	}
}
