package main

import "fmt"

type dataLoan struct {
	nik  string
	nama string
}

var allData = map[string]*dataLoan{}

func main() {

	a := new(dataLoan)
	a.nik = "1882183128"
	a.nama = "Mulia Ichsan"

	allData[a.nik] = a

	fmt.Println(allData)

	b := new(dataLoan)
	b.nik = "1882183128"
	b.nama = "Mulia Ichsan"

	allData[b.nik] = b

	fmt.Println(allData["1882183128"])
}
