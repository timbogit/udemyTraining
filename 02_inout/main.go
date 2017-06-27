package main

import "fmt"

func main() {
	fmt.Print("Enter degress in Fahrenheit: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)

	celsius := ((fahrenheit - 32) * 5 / 9)

	fmt.Println("Celsius is ", celsius)
}
