package main

import "fmt"

func canItems[T any](data []T) {
	for _, d := range data {
		fmt.Printf("%v ", d)
	}
	fmt.Println("")
}

func main() {
	canItems[int]([]int{1, 2, 4})
	canItems[string]([]string{"hello", "world"})
}
