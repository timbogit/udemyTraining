package main

import "fmt"

func main() {
	var lineBreak bool = false
	for i := 1; i < 101; i++ {
		if i % 3 == 0 {
			fmt.Print("Fizz")
			lineBreak = true
		}
		if i % 5  == 0 {
			fmt.Print("Buzz")
			lineBreak = true
		}
		if lineBreak {
			fmt.Println()
			lineBreak = false
		}
	}
}
