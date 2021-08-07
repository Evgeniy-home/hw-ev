package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var builtString strings.Builder
	var multiplier strings.Builder
	var previousChar rune
	var isEscaped bool
	for _, r := range s {
		if r == '\\' && !isEscaped {
			// если руна равна \\ и не бул
			isEscaped = true
			continue
		}
		if isEscaped {
			builtString.WriteRune(r)
			previousChar = r
			isEscaped = false
			continue
		}
		// Если r - цифра (0-9)
		if unicode.IsDigit(r) {
			// если предыдущая руна больше 0
			if previousChar > 0 {
				multiplier.WriteRune(r)
				// добавляем руну в строку
			} else {
				return "", nil
			}
		} else { // если руна не цифра
			if m, err := strconv.Atoi(multiplier.String()); err == nil && m > 0 && previousChar > 0 { // если строку нельзя конвертировать в инт, это не первый символ и предыдущая руна больше 0
				builtString.WriteString(strings.Repeat(string(previousChar), m-1)) // строим строку, где повторяем данный символ на один меньше, чем номер
				multiplier.Reset()                                                 // сбрасываем мультиплайер
			}
			builtString.WriteRune(r) //записываем руну, если предыдущие условия не сработали
			previousChar = r
		}
	}
	if m, err := strconv.Atoi(multiplier.String()); err == nil && m > 0 && previousChar > 0 {
		builtString.WriteString(strings.Repeat(string(previousChar), m-1))
	}
	if res := builtString.String(); len(res) > 0 {
		return builtString.String(), nil
	} else {
		return "", nil
	}
}
