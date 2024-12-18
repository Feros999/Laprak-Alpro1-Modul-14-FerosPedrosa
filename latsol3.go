package main

import "fmt"

func main() {
	var gelas1, gelas2, gelas3, gelas4 string
	var isAllCorrect bool

	isAllCorrect = true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)

		fmt.Scan(&gelas1, &gelas2, &gelas3, &gelas4)

		if gelas1 != "merah" || gelas2 != "kuning" || gelas3 != "hijau" || gelas4 != "ungu" {
			isAllCorrect = false
		}
	}

	fmt.Println("BERHASIL:", isAllCorrect)
}
