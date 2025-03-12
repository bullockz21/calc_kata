package main

import (
	"calculator/internal/delivery"
	"calculator/internal/usecase"
)

func main() {
	calc := usecase.NewCalculator()
	delivery.StartCLI(calc)
}
