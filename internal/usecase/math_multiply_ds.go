package usecase

import (
	"ryusuke410/golang-cli-tryout/pkg/cvalidator"
)

// MathMultiplyOutput represents the output for the multiply operation
type MathMultiplyOutput struct {
	Value string
}

// MathMultiplyInput represents the input for the multiply operation
type mathMultiplyInput struct {
	Numbers []float64 `validate:"min=2"`
	Base    int       `validate:"c_base"`
}

func NewMathMultiplyInput(numbers []float64, base int) (*mathMultiplyInput, error) {
	var err error
	input := &mathMultiplyInput{
		Numbers: numbers,
		Base:    base,
	}
	if err = cvalidator.Struct(input); err != nil {
		return nil, err
	}
	return input, nil
}
