package delivery

import (
	"bufio"
	"os"

	"calculator/internal/usecase"

	"github.com/sirupsen/logrus"
)

// StartCLI запускает консольный интерфейс для взаимодействия с пользователем.
func StartCLI(calc *usecase.Calculator) {
	scanner := bufio.NewScanner(os.Stdin)
	logrus.Info("Введите математическое выражение, разделяя оператор и числа пробелами:")

	for scanner.Scan() {
		input := scanner.Text()
		result, err := calc.Calculate(input)
		if err != nil {
			logrus.Error("Ошибка: ", err)
		} else {
			logrus.Info("Результат: ", result)
		}
	}
}
