package main

import "fmt"
import "math/rand"
import "sort"

/*Write a program in Go that will take a 5x5 matrix of random integers from 1 to 10 that
you will print out, then transpose and printout again in matrix form. Then take the matrix
and load it into a list and print it out in sorted order using the English words
for the numbers (i.e. 1 is “one”).*/

func main() {
	//declaring 2d slice
	matrix := createMatrix()
	printMatrix(matrix)
	transposed := transposeMatrix(matrix)
	printMatrix(transposed)
	list := convertMatrix(transposed)
	fmt.Println(list)
	sort.Ints(list)
	fmt.Println(list)

}

func createMatrix() [5][5]int {
	matrix := [5][5]int{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			randInt := rand.Intn(9) + 1 //makes random num s.t. 1<=num<=10
			matrix[i][j] = randInt
		}
	}
	return matrix
}

func transposeMatrix(inputMatrix [5][5]int) [5][5]int {
	outputMatrix := [5][5]int{}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			outputMatrix[i][j] = inputMatrix[j][i]
		}
	}
	return outputMatrix
}

func convertMatrix(matrix [5][5]int) []int {
	var list []int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			list = append(list, matrix[i][j])
		}
	}
	return list
}

func printMatrix(matrix [5][5]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("-------------")
}

func printWords(list []int) {
	strings := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

}
