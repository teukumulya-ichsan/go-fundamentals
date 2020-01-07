package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int32 = 10
	var y = int64(x) //converting int32 of x to int64
	fmt.Println(y)

	var z float64 = float64(y) //converting int64 of y to float64 save into z
	fmt.Println(z)

	var b float64 = 3.9
	var a int32 = int32(b)
	fmt.Println(a)

	var nilai string = "100"
	nilaiString, _ := strconv.Atoi(nilai) //method athoi return 2 values which is a value and error

	fmt.Println(nilaiString)

	var value int = 1000
	nilaiInt := strconv.Itoa(value) // sedangkan Itoa (convert to string return 1 value)

	fmt.Println(nilaiInt)

	//using parse (from string to string)
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseFloat("3.1415", 64))
	fmt.Println(strconv.ParseInt("-43", 10, 64))
	fmt.Println(strconv.ParseUint("34", 10, 64))

	//using format (from typedata all to other string)

	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatFloat(3.1415, 'E', -1, 64))
	fmt.Println(strconv.FormatInt(-42, 16))
	fmt.Println(strconv.FormatUint(34, 16))
}
