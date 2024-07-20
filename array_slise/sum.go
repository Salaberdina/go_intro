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

func SumAllTails(slises ...[]int) []int {
	result := []int{}
	for _, s := range slises {
		if len(s) < 2 {
			result = append(result, 0)
		} else {
			result = append(result, SumSlise(s[1:]))
		}
	}
	return result
}
