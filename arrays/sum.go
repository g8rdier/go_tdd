package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthofNumbers := len(numbersToSum)
	sums := make([]int, lengthofNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return nil
}
