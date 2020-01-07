package main

import (
	"fmt"
	"runtime"
)

func main() {

	for i := 1; i <= 30; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FooBar")
		case i%3 == 0:
			fmt.Println("Foo")
		case i%5 == 0:
			fmt.Println("Bar")
		default:
			fmt.Println(i)
		}
	}

	sistemOperasi := runtime.GOOS

	switch sistemOperasi {
	case "darwin":
		fmt.Println("Mac OS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Unknown")
	}
}
