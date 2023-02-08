package main

import "fmt"

func area(lenght, width int) int {
	ar := lenght * width
	return ar
}

func main() {
	fmt.Printf("area of rectangle is : %d", area(12, 10))
}
