package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	doubleFunc := createTransformerFunction(2)
	tripleFunc := createTransformerFunction(3)

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(transformed)

	double := transformNumbers(&numbers, doubleFunc)
	triple := transformNumbers(&numbers, tripleFunc)

	fmt.Println(double)
	fmt.Println(triple)

}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformerFunction(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
