package usecase

import (
	"context"

	"ryusuke410/golang-cli-tryout/internal/domain/apperrors"
)

func (u *mathUseCase) Multiply(ctx context.Context, input *MathMultiplyInput) (*MathMultiplyOutput, error) {
	product, err := u.mathService.Product(input.numbers)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	valuestr, err := u.mathService.FormatNumber(product, input.base)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	return &MathMultiplyOutput{Value: valuestr}, nil
}
