package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)

	if bilangan <= 1 {
		fmt.Println("bukan prima")
	} else {
		prima := true
		for i := 2; i*i <= bilangan; i++ {
			if bilangan%i == 0 {
				prima = false
				break
			}
		}

		if prima {
			fmt.Println("prima")
		} else {
			fmt.Println("bukan prima")
		}
	}
}
