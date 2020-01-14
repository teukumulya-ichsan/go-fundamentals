package main

import "fmt"

func main() {

	// looping jika i habis dibagi 2 (genap) akan di skip untuk tampilkan
	// pengecualiannya sampai i lebih kecil dari 8 maka looping di stop
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		if i > 8 {
			break
		}
		fmt.Println("Angka", i)
	}
}
