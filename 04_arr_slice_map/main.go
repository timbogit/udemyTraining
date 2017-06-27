package main

import "fmt"

func main()  {
	var intArr = make([]int, 3, 9)
	x := [6]string{"a", "b", "c", "d", "e", "f",}
	fmt.Println("Length of intArr: ", intArr)
	fmt.Println("x[2:5] = ", x[2:5])

	y := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}

	fmt.Println("smallest int in y = ", smallest(y))
}

func smallest(x []int) int {
	smallestIndex := 0
	for idx, val := range x {
		if x[smallestIndex] > val {
			smallestIndex = idx
		}
	}
	return x[smallestIndex]
}
