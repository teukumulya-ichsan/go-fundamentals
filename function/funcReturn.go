package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	var randomValue int

	randomValue = randowWithRange(1, 10)

	fmt.Println("random number:", randomValue)
}

func randowWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min

	return value

}
