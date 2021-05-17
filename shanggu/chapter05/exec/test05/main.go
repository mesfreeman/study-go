package main

import "fmt"

func main() {
	matrix := [][]int{
		{-5},
	}
	fmt.Println(len(matrix), findNumberIn2DArray(matrix, -5))
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	for k1 := 0; k1 < len(matrix); k1++ {
		for k2 := len(matrix[0]); k2 > 0; k2-- {
			if matrix[k1][k2-1] == target {
				return true
			}
			if matrix[k1][k2-1] < target {
				goto continuek1
			}
		}
	continuek1:
	}

	return false
}
