package usecase

import "fmt"

// MathAddOutput represents the output for the add operation
type MathAddOutput struct {
	Value float64
}

// mathAddInput represents the input for the add operation
type mathAddInput struct {
	Numbers []float64
	Round   bool
}

func NewMathAddInput(numbers []float64, round bool) (*mathAddInput, error) {
	if len(numbers) < 2 {
		return nil, fmt.Errorf("at least 2 numbers are required")
	}
	return &mathAddInput{
		Numbers: numbers,
		Round:   round,
	}, nil
}
