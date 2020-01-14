package main

import "fmt"

func main() {

	number := []float64{1, 2, 3, 4}

	fmt.Println("sum:", getSum(number))
	fmt.Println("average:", getAverage(number))
}

// func getSum(a int, b int) int {
// 	return a + b
// }

func getSum(data []float64) float64 {
	total := 0.0
	for _, value := range data {
		total += value
	}
	return total
}

func getAverage(data []float64) float64 {
	total := 0.0

	for _, value := range data {
		total += value
	}

	return total / float64(len(data))
}
