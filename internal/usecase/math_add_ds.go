package usecase

import (
	"ryusuke410/golang-cli-tryout/pkg/cvalidator"
)

// MathAddOutput represents the output for the add operation
type MathAddOutput struct {
	Value float64
}

// mathAddInput represents the input for the add operation
type mathAddInput struct {
	Numbers []float64 `validate:"min=2"`
	Round   bool
}

func NewMathAddInput(numbers []float64, round bool) (*mathAddInput, error) {
	var err error
	input := &mathAddInput{
		Numbers: numbers,
		Round:   round,
	}
	if err = cvalidator.Struct(input); err != nil {
		return nil, err
	}
	return input, nil
}
