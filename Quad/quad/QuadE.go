package piscine

import "fmt"

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 && col == 0 {
				fmt.Print("A")
			} else if row == 0 && col == x-1 {
				fmt.Print("C")
			} else if row == y-1 && col == 0 {
				fmt.Print("C")
			} else if row == y-1 && col == x-1 {
				fmt.Print("A")
			} else if row == 0 || row == y-1 {
				fmt.Print("B")
			} else if col == 0 || col == x-1 {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
