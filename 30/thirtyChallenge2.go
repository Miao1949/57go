package main

import "fmt"

func main() {
	for row := -1; row <= 12; row++ {
		for col := -1; col <= 12; col++ {
			// The empty square
			if row == -1 && col == -1 {
				fmt.Print("\t")
			} else if row == -1 {
				if col < 12 {
					fmt.Printf("%d\t", col)
				} else {
					fmt.Printf("%d\n", col)
				}
			} else if col == -1 {
				fmt.Printf("%d\t", row)
			} else if col == 12 {
				// Last row. Print newline.
				fmt.Printf("%d\n", row*col)
			} else {
				fmt.Printf("%d\t", row*col)
			}
		}
	}
}
