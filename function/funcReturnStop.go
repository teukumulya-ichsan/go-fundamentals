package main

import "fmt"

func main() {
	divideNumber(10, 0)
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / % d =  %d\n", m, n, res)
}
