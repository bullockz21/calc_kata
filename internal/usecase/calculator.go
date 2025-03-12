package usecase

import (
	"errors"
	"fmt"
	"strings"

	"calculator/internal/domain"
)

// Calculator представляет бизнес-логику калькулятора.
type Calculator struct{}

// NewCalculator создаёт новый экземпляр Calculator.
func NewCalculator() *Calculator {
	return &Calculator{}
}

// ParseExpression анализирует входную строку и возвращает объект, реализующий domain.Expression.
func (c *Calculator) ParseExpression(input string) (domain.Expression, error) {
	tokens := strings.Fields(input)
	if len(tokens) != 3 {
		return nil, errors.New("некорректное количество аргументов; формат: [число/римское] [оператор] [число/римское]")
	}

	operand1, err := domain.ParseOperand(tokens[0])
	if err != nil {
		return nil, fmt.Errorf("ошибка при разборе первого операнда: %w", err)
	}

	operand2, err := domain.ParseOperand(tokens[2])
	if err != nil {
		return nil, fmt.Errorf("ошибка при разборе второго операнда: %w", err)
	}

	operator := tokens[1]
	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		return nil, errors.New("некорректный оператор; допустимы: +, -, *, /")
	}

	return domain.BasicExpression{
		A:        operand1,
		B:        operand2,
		Operator: operator,
	}, nil
}

// Calculate парсит выражение и вычисляет результат.
func (c *Calculator) Calculate(input string) (int, error) {
	expr, err := c.ParseExpression(input)
	if err != nil {
		return 0, err
	}
	return expr.Evaluate()
}
