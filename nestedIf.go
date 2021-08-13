// Sebuah wahana permainan mengenakan tarif Rp15.000 untuk satu kali bermain.
// Selain itu juga disyaratkan pengunjung berusia minimal 7 tahun dan tinggi badan minimal 125 cm.
// Khusus untuk anak dibawah usia 10 tahun wajib didampingi oleh orang dewasa.

// INSTRUCTIONS
// 1. Jika uang tidak cukup, tampilkan "Maaf, uang kamu tidak cukup".
//    Jika uang cukup, maka lanjut ke step berikutnya.
// 2. Jika usia pengunjung kurang dari 7 tahun atau tinggi badan kurang dari 125, tampilkan
//    "Kamu belum diperbolehkan menaiki wahana ini".
//    Jika tidak, maka lanjut ke step berikutnya.
// 3. Jika pengunjung berusia dibawah 10 tahun dan tidak didampingi oleh orang dewasa, tampilkan
//    "Kamu harus ditemani oleh orang dewasa untuk menaiki wahana ini"
//    Jika tidak, maka tampilkan "Selamat menikmati wahana ini".

package main

import "fmt"

func main() {
	money := 15000
	age := 9
	high := 159

	if money < 15000 {
		fmt.Println("Maaf, uang kamu tidak cukup")
	} else if age < 7 || high < 150 {
		fmt.Println("Kamu belum diperbolehkan menaiki wahana ini")
	} else if age < 10 {
		fmt.Println("Kamu harus ditemani oelh orang dewasa untuk menaiki wahana ini")
	} else {
		fmt.Println("Selamat menikmati wahana ini")
	}
}
