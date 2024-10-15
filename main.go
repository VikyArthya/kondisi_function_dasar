package main

import (
	"fmt"
)

// //function untuk menghitung luas dan keliling persegi
// func hitungPersegi(sisi float64) (float64, float64) {
// 	luas := sisi * sisi
// 	keliling := 4 * sisi
// 	return luas, keliling
// }

// func main() {
// 	var sisi float64

// 	fmt.Print("Masukkan panjang sisi persegi: ")
// 	fmt.Scanln(&sisi)

// 	luas, keliling := hitungPersegi(sisi)

// 	fmt.Printf("Luas persegi: %.2f\n", luas)
// 	fmt.Printf("Keliling persegi: %.2f\n", keliling)
// }

// //Function harga setelah diskon
// func hargaSetelahDiskon(harga []float64, diskon float64) []float64 {
// 	var hargaDiskon []float64
// 	for _, h := range harga {
// 		hargaSetelah := h - (h * diskon / 100)
// 		hargaDiskon = append(hargaDiskon, hargaSetelah)
// 	}
// 	return hargaDiskon
// }

// func main() {
// 	hargaBarang := []float64{100000, 200000, 150000, 50000}
// 	diskon := 10.0 // diskon 10%

// 	hasilDiskon := hargaSetelahDiskon(hargaBarang, diskon)

// 	fmt.Println("Harga setelah diskon:", hasilDiskon)
// }

// func hitungDiskon(harga []float64) []float64 {
// 	var diskon float64
// 	jumlahBarang := len(harga)

// 	if jumlahBarang >= 2 && jumlahBarang < 4 {
// 		diskon = 10.0
// 	} else if jumlahBarang == 4 {
// 		diskon = 50.0
// 	} else if jumlahBarang > 4 {
// 		diskon = 75.0
// 	} else {
// 		diskon = 0.0
// 	}

// 	var hargaDiskon []float64
// 	for _, h := range harga {
// 		hargaSetelah := h - (h * diskon / 100)
// 		hargaDiskon = append(hargaDiskon, hargaSetelah)
// 	}
// 	return hargaDiskon
// }

// func main() {
// 	hargaBarang := []float64{100000, 200000, 150000, 50000, 300000}
// 	hasilDiskon := hitungDiskon(hargaBarang)
// 	fmt.Println("Harga setelah diskon:", hasilDiskon)
// }

// // hitunglah total gaji seorang karyawan jika 1 jamnya dibayar 50rb, kemudian lemburnya 1 jamnya 60rb
// // total dia bekerja 40 jam, leburnya 5 jam, buatkan 1 fuction untuk menghitung total gaji dia

// func totalGaji(jamKerja int, jamLembur int) int {
// 	var hargaJamkerja int = 50000
// 	var hargaJamlembur int = 60000
// 	totalGajikaryawan := (jamKerja * hargaJamkerja) + (jamLembur * hargaJamlembur)
// 	return totalGajikaryawan
// }

// func main() {
// 	var jamKerja int = 45
// 	var jamLembur int = 10
// 	result := totalGaji(jamKerja, jamLembur)
// 	fmt.Println(result)
// }

var (
	adidasPrice int = 200000
	pumaPrice   int = 150000
	kappaPrice  int = 600000
)

func hitungTotal(adidas, puma, kappa bool) int {
	var totalHarga, diskon int
	if adidas {
		totalHarga += adidasPrice
	}
	if puma {
		totalHarga += pumaPrice
	}
	if kappa {
		totalHarga += kappaPrice
	}

	if adidas && puma {
		diskon = 50000
	} else if puma && kappa {
		diskon = 150000
	} else if adidas && kappa {
		diskon = 75000
	}

	totalHarga -= diskon

	return totalHarga
}

func main() {
	total := hitungTotal(true, true, false)
	fmt.Println("Total harga setelah diskon:", total)

	total = hitungTotal(false, true, true)
	fmt.Println("Total harga setelah diskon:", total)

	total = hitungTotal(true, false, true)
	fmt.Println("Total harga setelah diskon:", total)

	total = hitungTotal(false, true, false)
	fmt.Println("Total harga setelah diskon:", total)
}
