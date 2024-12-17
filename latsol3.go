package main

import (
	"fmt"
)

func main() {
	// Definisikan urutan warna yang diharapkan
	expectedColors := []string{"merah", "kuning", "hijau", "ungu"}
	var results []bool

	// Lakukan 5 percobaan
	for i := 0; i < 5; i++ {
		var colors [4]string
		fmt.Printf("Percobaan ke-%d:\n", i+1)

		// Minta input warna dari pengguna
		for j := 0; j < 4; j++ {
			fmt.Printf("Masukkan warna untuk tabung reaksi %d: ", j+1)
			fmt.Scan(&colors[j])
		}

		// Periksa apakah urutan warna sesuai
		isSuccess := true
		for k := 0; k < 4; k++ {
			if colors[k] != expectedColors[k] {
				isSuccess = false
				break
			}
		}

		// Simpan hasil percobaan
		results = append(results, isSuccess)
	}

	// Tampilkan hasil percobaan
	fmt.Println("\nHasil percobaan:")
	for i, result := range results {
		fmt.Printf("Percobaan ke-%d: %t\n", i+1, result)
	}
}
