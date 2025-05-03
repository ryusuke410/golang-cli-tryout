package usecase

import (
	"ryusuke410/golang-cli-tryout/internal/pkg/cvalidator"
)

type MathAddOutput struct {
	Value float64
}

type MathAddInput struct {
	numbers []float64 `validate:"min=2"`
	round   bool
}

func NewMathAddInput(numbers []float64, round bool) (*MathAddInput, error) {
	var err error
	input := &MathAddInput{
		numbers: numbers,
		round:   round,
	}
	if err = cvalidator.Struct(input); err != nil {
		return nil, err
	}
	return input, nil
}
