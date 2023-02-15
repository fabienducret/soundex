package soundex

import (
	"strings"
)

const MAX_LENGTH = 4
const NOT_A_DIGIT = ""

func Encode(word string) string {
	return zeroPad(head(upper(word)) + encodedDigits(tail(word)))
}

func zeroPad(word string) string {
	zeroNeeded := MAX_LENGTH - len(word)
	return word + strings.Repeat("0", zeroNeeded)
}

func head(word string) string {
	return word[0:1]
}

func upper(word string) string {
	return strings.ToUpper(word)
}

func encodedDigits(word string) string {
	var encoding string

	for _, letter := range word {
		if isComplete(encoding) {
			break
		}

		digit := encodedDigit(string(letter))
		if digit != NOT_A_DIGIT && digit != lastDigit(encoding) {
			encoding += digit
		}
	}

	return encoding
}

func isComplete(encoding string) bool {
	return len(encoding) == MAX_LENGTH-1
}

func tail(word string) string {
	return word[1:]
}

func encodedDigit(letter string) string {
	encodings := map[string]string{
		"b": "1", "f": "1", "p": "1", "v": "1",
		"c": "2", "g": "2", "j": "2", "k": "2", "q": "2", "s": "2", "x": "2", "z": "2",
		"d": "3", "t": "3",
		"l": "4",
		"m": "5", "n": "5",
		"r": "6",
	}

	value, ok := encodings[lower(letter)]

	if ok {
		return value
	}

	return NOT_A_DIGIT
}

func lastDigit(encoding string) string {
	if len(encoding) == 0 {
		return NOT_A_DIGIT
	}

	last := encoding[len(encoding)-1:]

	return last
}

func lower(word string) string {
	return strings.ToLower(word)
}
