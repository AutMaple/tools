package str

import (
	"unicode"
)

func PascalToCamel(word string) string {
	runes := []rune(word)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}
