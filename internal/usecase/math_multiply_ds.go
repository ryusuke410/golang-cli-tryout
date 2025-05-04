package usecase

import (
	"ryusuke410/golang-cli-tryout/internal/pkg/cvalidator"
)

type MathMultiplyOutput struct {
	Value string
}

type MathMultiplyInput struct {
	numbers []float64 `validate:"min=2"`
	base    int       `validate:"c_base"`
}

func NewMathMultiplyInput(numbers []float64, base int) (*MathMultiplyInput, error) {
	var err error
	input := &MathMultiplyInput{
		numbers: numbers,
		base:    base,
	}
	if err = cvalidator.Struct(input); err != nil {
		return nil, err
	}
	return input, nil
}
