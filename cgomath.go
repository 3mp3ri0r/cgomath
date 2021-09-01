package cgomath

func sum(numbers []int, result int) int {
	if len(numbers) == 0 {
		return result
	}

	number := numbers[0]
	return sum(numbers[1:], result+number)
}

func Sum(numbers []int) int {
	return sum(numbers, 0)
}
