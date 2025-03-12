package usecase

import (
	"testing"
)

func TestCalculator_Calculate(t *testing.T) {
	calc := NewCalculator()

	tests := []struct {
		input    string
		expected int
		err      bool
	}{
		{"1 + 2", 3, false},
		{"V * III", 15, false}, // V = 5, III = 3, 5*3=15
		{"10 / 2", 5, false},
		{"10 / 0", 0, true}, // деление на ноль
		{"1 ^ 2", 0, true},  // недопустимый оператор
		{"1 +", 0, true},    // недостаточное число аргументов
	}

	for _, tt := range tests {
		result, err := calc.Calculate(tt.input)
		if tt.err && err == nil {
			t.Errorf("ожидалась ошибка для входа '%s', но ошибки не произошло", tt.input)
		} else if !tt.err && err != nil {
			t.Errorf("не ожидалась ошибка для входа '%s', но получили: %v", tt.input, err)
		} else if !tt.err && result != tt.expected {
			t.Errorf("для входа '%s' ожидали результат %d, получено %d", tt.input, tt.expected, result)
		}
	}
}
