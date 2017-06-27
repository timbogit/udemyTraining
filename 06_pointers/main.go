package main

import "fmt"

func swap(x, y *int) {
	*x ^= *y
	*y ^= *x
	*x ^= *y
}

func main() {
	x := 12
	y := 34
	swap(&x, &y)
	fmt.Println("x is now ", x, ", and y is now ", y)
}