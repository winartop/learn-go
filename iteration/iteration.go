package iteration

import "strings"

// func Repeat(Character string) string {
// 	return "aaaaa"
// }

const repeatedCount = 5

func Repeat(Character string) string {
	// var repeated string
	var repeated strings.Builder

	for i := 0; i < repeatedCount; i++ {
		// repeated += Character
		repeated.WriteString(Character)
	}

	return repeated.String()
}
