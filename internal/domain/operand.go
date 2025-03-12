package domain

import (
	"errors"
	"strconv"
)

// AllowedRoman содержит допустимые римские числа и их соответствия арабским числам.
var AllowedRoman = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

// ParseOperand преобразует строковое значение в число, проверяя, является ли оно римским или арабским.
func ParseOperand(token string) (int, error) {
	if value, ok := AllowedRoman[token]; ok {
		return value, nil
	}

	num, err := strconv.Atoi(token)
	if err != nil {
		return 0, errors.New("некорректное число: " + token)
	}
	if num < 1 || num > 10 {
		return 0, errors.New("число должно быть в диапазоне от 1 до 10")
	}
	return num, nil
}
