package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Unpack принимает строку и возвращает распакованную строку.
func Unpack(s string) (string, error) {
	var result strings.Builder
	length := len(s)
	i := 0

	for i < length {
		char := s[i]
		i++

		if char == '\\' { // Проверка escape-последовательности
			if i < length {
				result.WriteByte(s[i])
				i++
			} else {
				return "", errors.New("некорректная строка")
			}
			continue
		}

		if i < length && unicode.IsDigit(rune(s[i])) { // Если следующий символ - цифра
			countStr := string(s[i])
			i++

			// Если следующая цифра
			for i < length && unicode.IsDigit(rune(s[i])) {
				countStr += string(s[i])
				i++
			}

			count, err := strconv.Atoi(countStr)
			if err != nil {
				return "", errors.New("некорректная строка")
			}

			result.WriteString(strings.Repeat(string(char), count))
		} else {
			result.WriteByte(char)
		}
	}

	return result.String(), nil
}

func main() {
	testCases := []string{
		"a4bc2d5e",  // "aaaabccddddde"
		"abcd",      // "abcd"
		"45",        // ""
		"",          // ""
		"qwe\\4\\5", // "qwe45"
		"qwe\\45",   // "qwe44444"
		"qwe\\\\5",  // "qwe\\\\\\"
	}

	for _, testCase := range testCases {
		result, err := Unpack(testCase)
		if err != nil {
			fmt.Printf("Input: %q, Error: %s\n", testCase, err)
		} else {
			fmt.Printf("Input: %q, Output: %q\n", testCase, result)
		}
	}
}
