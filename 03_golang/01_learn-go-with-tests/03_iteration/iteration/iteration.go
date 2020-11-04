package iteration

const RepeatCount = 5

// Repeat take in a string and returns a 5 time rpeat of the string
func Repeat(a string, count int) string{
	output := ""
	if count < 0 {
		count = RepeatCount
	}
	for i :=0;i < count; i++ {
		output+=a
	}
	return output
}
