package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

func ExampleReapeat(character string, number int) string {
	var repeated strings.Builder
	for i := 0; i < number; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
