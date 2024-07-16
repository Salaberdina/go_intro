package iteretions

func Repeat(a string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result = result + a
	}
	return result
}
