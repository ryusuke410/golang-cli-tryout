package usecase

import "context"

func (u *mathUseCase) Multiply(ctx context.Context, input *mathMultiplyInput) (*MathMultiplyOutput, error) {
	product := u.mathService.Product(input.Numbers)
	return &MathMultiplyOutput{Value: u.mathService.FormatNumber(product, input.Base)}, nil
}
