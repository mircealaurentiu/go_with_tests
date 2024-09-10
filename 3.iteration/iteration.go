package iteration

func Repeat(char string, times int) string {
	return_string := ""
	for i := 0; i < times; i++ {
		return_string += char
	}
	return return_string
}
