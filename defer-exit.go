package main

import (
	"fmt"
)

// defer digunakan untuk mengakhirkan eksekusi sebuah statement tepat sebelum blok fungsi selesai.
func orderSomeFood(menu string) {
	defer fmt.Println("Terima kasih, silahkan tunggu")

	if menu == "Mie gacoan lvl 8" {
		fmt.Print("\n", "Pilihan yang bagus 😋, ")
		fmt.Print("Mie gacoan lvl 8 ditempat kami paling enak!", "\n")
		return
	}

	fmt.Println("Pesanan anda: ", menu)
}

// kombinasi defer dan IIFE
func nomerAntrian() {
	number := 3

	if number == 3 {
		fmt.Println("\n", "antrian ke-1")
		func() {
			defer fmt.Println(" antrian ke-3")
		}()
	}

	fmt.Println(" antrian ke-2")
}

/* exit digunakan untuk menghentikan program secara paksa,
(ingat, menghentikan program, tidak seperti return yang hanya menghentikan blok kode).
*/

// func keluar() {
// 	defer fmt.Println("Gimana kabarnya?")
// 	os.Exit(1)
// 	fmt.Println("\nSup! guys")
// }
