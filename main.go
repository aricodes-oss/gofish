package main

import "fmt"

func main() {
	x := make([]int, 32)
	for idx := range x {
		x[idx] = idx
	}

	fmt.Println(x[1:3])
}
