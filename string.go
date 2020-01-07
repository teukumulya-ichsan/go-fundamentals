package main

import "fmt"

func main() {
	fmt.Println("Belajar Go-Lang")
	fmt.Println("Belajar" + " " + "GO")
	fmt.Println(len("Belajar")) // return 7
	fmt.Println("Belajar"[0])   // return byte of B not return B
	fmt.Println("Belajar"[0:4]) // should return Bela
	fmt.Println("Belajar"[2:])  // should return from index 2 until end = lajar
	fmt.Println("Belajar"[:4])  // shhould return Bela
}
