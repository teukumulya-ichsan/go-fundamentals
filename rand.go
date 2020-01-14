package main

import (
	"fmt"
	"time"
)

func main() {

	currenTime := time.Now()
	timeFormat := currenTime.Format("20060102150405")

	fmt.Println(timeFormat)
}
