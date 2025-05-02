package usecase

import (
	"context"
)

// TransformUseCase represents the transform use case
type ITransformUseCase interface {
	Base64Encode(ctx context.Context, input *transformBase64Input) (*TransformBase64Output, error)
	Escape(ctx context.Context, input *transformEscapeInput) (*TransformEscapeOutput, error)
	UrlEncode(ctx context.Context, input *transformUrlEncodeInput) (*TransformUrlEncodeOutput, error)
}

// transformUseCase implements TransformUseCase
type transformUseCase struct{}

var _ ITransformUseCase = (*transformUseCase)(nil)

func NewTransformUseCase() *transformUseCase {
	return &transformUseCase{}
}
