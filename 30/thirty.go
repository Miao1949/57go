package main

import "fmt"

func main() {
	for i1 :=0; i1 <= 12; i1++ {
		for i2 :=0; i2 <= 12; i2++ {
			fmt.Printf("%d x %d = %d\n", i1, i2, i1 * i2)
		}
	}
}
