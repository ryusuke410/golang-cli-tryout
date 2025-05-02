package usecase

import "fmt"

// MathMultiplyOutput represents the output for the multiply operation
type MathMultiplyOutput struct {
	Value string
}

// MathMultiplyInput represents the input for the multiply operation
type mathMultiplyInput struct {
	Numbers []float64
	Base    int
}

func NewMathMultiplyInput(numbers []float64, base int) (*mathMultiplyInput, error) {
	if len(numbers) < 2 {
		return nil, fmt.Errorf("at least 2 numbers are required")
	}
	return &mathMultiplyInput{
		Numbers: numbers,
		Base:    base,
	}, nil
}
