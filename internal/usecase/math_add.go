package usecase

import "context"

func (u *mathUseCase) Add(ctx context.Context, input *MathAddInput) (*MathAddOutput, error) {
	sum := u.mathService.Sum(input.numbers)
	if input.round {
		sum = u.mathService.Round(sum)
	}
	return &MathAddOutput{Value: sum}, nil
}
