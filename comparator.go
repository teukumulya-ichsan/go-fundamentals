package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int32 = 1
  var b int32 = 1

	//compare equal
	resultEqual := a == b // true

	fmt.Println("resultEqual : " + strconv.FormatBool(resultEqual)) // return true cause a is equal b

	resultNotEqual := a != b
	fmt.Println("resulNotEqual : " + strconv.FormatBool(resultNotEqual)) // return false because a is equal b

	resultLower := a < b
	fmt.Println("resultLower : " + strconv.FormatBool(resultLower)) //return false cause a is not lower b

	resultLowerThen := a <= b
	fmt.Println("resultLowerThen : " + strconv.FormatBool(resultLowerThen)) //return true cause a is lower equal b

	resultGreater := a > b
	fmt.Println("resultGreater : " + strconv.FormatBool(resultGreater)) //return false cause a is not greater b

	resultGreaterThen := a >= b
	fmt.Println("resultGreaterThen : " + strconv.FormatBool(resultGreaterThen)) //return true cause a is greater equal b

}
