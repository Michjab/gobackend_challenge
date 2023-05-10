package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// Function checks if there was no overflow during integer multiplication for two integers (x,y) after sumed (d). Returns boolean if overflow
func isOverflowMultiply(d, x, y int) bool {
	return d/y != x
}

// Function checks if two given integers can overflow integer (MaxInt, MinInt) value after adding (x,y). Returns boolean if overflow
func isOverflowAdd(x, y int) bool {
	if y > 0 {
		if x > math.MaxInt32-y {
			return true
		}
	} else {
		if x < math.MinInt32-y {
			return true
		}
	}
	return false
}

// Function sums given []string (record). Returns int value (sum) or an error
func sumRow(record []string) (int, error) {
	recordSum := 0

	for index := range record {
		value, err := strconv.Atoi(record[index])
		if err != nil {
			ErrNotANumber := errors.New("csv contain value that is not an integer")
			return 0, ErrNotANumber
		}
		if isOverflowAdd(recordSum, value) {
			ErrOverflow := errors.New("integer adding overflow")
			return 0, ErrOverflow
		}
		recordSum += value
	}
	return recordSum, nil
}

// Function converts given []string (record). Returns []integer (intSlice) or an error
func strToInt(record []string) ([]int, error) {
	intSlice := make([]int, len(record))
	for index := range record {
		value, err := strconv.Atoi(record[index])
		if err != nil {
			return nil, fmt.Errorf("csv should contain integer values only: %w", err)
		}
		intSlice[index] = value
	}
	return intSlice, nil
}

// Function multiplies given []integer (records). Returns multiplied integers (result) or an error
func multiply(record []int) (int, error) {
	result := 1
	for _, value := range record {
		if value == 0 { //in case value is zero everything is be zero
			result = 0
			return result, nil
		}
		resultTmp := result * value
		if isOverflowMultiply(resultTmp, result, value) {
			ErrOverflow := errors.New("integer multiplication overflow")
			return 0, ErrOverflow
		}
		result = resultTmp
	}
	return result, nil
}

// Function transpose input [][]string (inM). Columns are rows and rows are columns. Return [][]string (transposedM) or an error
func transpose(inM [][]string) ([][]string, error) {
	matrixLen := len(inM[0])
	transposedM := make([][]string, matrixLen)

	for i := range transposedM {
		transposedM[i] = make([]string, matrixLen)
	}
	for i := 0; i < matrixLen; i++ {
		for j := 0; j < matrixLen; j++ {
			_, err := strconv.Atoi(inM[j][i])
			if err != nil {
				ErrNotANumber := errors.New("csv contain value that is not an integer")
				return nil, ErrNotANumber
			}
			transposedM[i][j] = inM[j][i]
		}
	}
	return transposedM, nil
}
