package main

import "fmt"

func main() {
	x := make([]int, 32)
	for idx := range x {
		x[idx] = idx
	}

	l := len(x)

	fmt.Println(x[:l])
	fmt.Println('\n')
}
