package usecase

import (
	"context"

	"ryusuke410/golang-cli-tryout/internal/domain/service"
)

// IMathUseCase represents multiple use cases
type IMathUseCase interface {
	Add(ctx context.Context, input *MathAddInput) (*MathAddOutput, error)
	Multiply(ctx context.Context, input *MathMultiplyInput) (*MathMultiplyOutput, error)
}

// mathUseCase implements IMathUseCase
type mathUseCase struct {
	mathService service.IMathService
}

var _ IMathUseCase = (*mathUseCase)(nil)

func NewMathUseCase(mathService service.IMathService) *mathUseCase {
	return &mathUseCase{
		mathService: mathService,
	}
}
