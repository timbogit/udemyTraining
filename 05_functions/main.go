package main

import "fmt"

func main()  {

	fmt.Println("Fibonacci of 5 = ", fib(5) )

	y := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}

	fmt.Println("greatest int in y = ", greatest(y...))
	fmt.Println("sum of y = ", sum(y))

	nexOdd := makeOddGenerator()
	fmt.Println(nexOdd()) // 1
	fmt.Println(nexOdd()) // 3
	fmt.Println(nexOdd()) // 5
}

func greatest(x ...int) int {
	greatestIndex := 0
	for idx, val := range x {
		if x[greatestIndex] < val {
			greatestIndex = idx
		}
	}
	return x[greatestIndex]
}

func fib(n uint) uint {
	switch n {
	case 0, 1: return uint(1)
	default: return fib(n - 1) + fib (n - 2)
	}
}

func sum(x []int) int {
	var total int
	for _, val := range x {
		total += val
	}
	return total
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}