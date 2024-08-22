package main

import "fmt"

func main() {

	sum := sumup(1, 15, 58)

	fmt.Println(sum)
}

func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
