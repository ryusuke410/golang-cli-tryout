package usecase

import (
	"context"

	"ryusuke410/golang-cli-tryout/internal/domain/apperrors"
)

func (u *mathUseCase) Multiply(ctx context.Context, input *mathMultiplyInput) (*MathMultiplyOutput, error) {
	product, err := u.mathService.Product(input.Numbers)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	return &MathMultiplyOutput{Value: u.mathService.FormatNumber(product, input.Base)}, nil
}
