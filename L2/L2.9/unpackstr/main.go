package unpackstr

import (
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(s string) (string, error) {

	var sb strings.Builder

	runes := []rune(s)

	for i, c := range runes {

		if count, ok := DoIT(runes, c, i); ok {
			for range count - 1 {
				sb.WriteRune(runes[i-1])
			}

		} else {
			if c != '\\' {
				sb.WriteRune(c)
			}
		}

	}

	if _, err := strconv.Atoi(sb.String()); err == nil {
		return "", os.ErrInvalid
	} else {
		return sb.String(), nil
	}
}

func DoIT(runes []rune, char rune, i int) (int, bool) {

	if i-1 < 0 {
		return 0, false
	}

	prevRune := runes[i-1]

	count, err := strconv.Atoi(string(char))

	if err != nil {
		return 0, false
	}

	if unicode.IsLetter(prevRune) {
		return count, true
	}

	if i-2 < 0 {
		return 0, false
	}

	prevprevRune := runes[i-2]

	if prevprevRune == '\\' {
		return count, true
	}

	return 0, false

}
