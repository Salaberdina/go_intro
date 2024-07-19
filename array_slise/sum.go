package arrayslise

func Sum(numbers [5]int) int {
	var result int
	for i := range numbers {
		result = result + numbers[i]
	}
	return result
}

func SumSlise(numbers []int) int {
	var result int
	for i := range numbers {
		result = result + numbers[i]
	}
	return result
}

func SumAll(slises ...[]int) []int {
	result := []int{}
	for i := range slises {
		sum := SumSlise(slises[i])
		result = append(result, sum)
	}
	return result
}
