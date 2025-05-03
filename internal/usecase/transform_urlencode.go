package usecase

import (
	"context"
	"strings"
)

func (u *transformUseCase) UrlEncode(ctx context.Context, input *TransformUrlEncodeInput) (*TransformUrlEncodeOutput, error) {
	return &TransformUrlEncodeOutput{
		Value: strings.ReplaceAll(strings.ReplaceAll(input.text, " ", "%20"), "&", "%26"),
	}, nil
}
