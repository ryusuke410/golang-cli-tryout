package usecase

import "context"

func (u *mathUseCase) Add(ctx context.Context, input *mathAddInput) (*MathAddOutput, error) {
	sum := u.mathService.Sum(input.Numbers)
	if input.Round {
		sum = u.mathService.Round(sum)
	}
	return &MathAddOutput{Value: sum}, nil
}
