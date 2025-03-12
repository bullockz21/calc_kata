package domain

import "errors"

// Expression определяет интерфейс для математических выражений.
type Expression interface {
	Evaluate() (int, error)
}

// BasicExpression реализует Expression и хранит два операнда и оператор.
type BasicExpression struct {
	A, B     int
	Operator string
}

// Evaluate выполняет арифметическую операцию над операндами BasicExpression.
func (be BasicExpression) Evaluate() (int, error) {
	switch be.Operator {
	case "+":
		return be.A + be.B, nil
	case "-":
		return be.A - be.B, nil
	case "*":
		return be.A * be.B, nil
	case "/":
		if be.B == 0 {
			return 0, errors.New("деление на ноль")
		}
		return be.A / be.B, nil
	default:
		return 0, errors.New("неподдерживаемый оператор")
	}
}
