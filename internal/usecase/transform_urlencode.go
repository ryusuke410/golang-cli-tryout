package usecase

import (
	"context"
	"strings"
)

func (u *transformUseCase) UrlEncode(ctx context.Context, input *transformUrlEncodeInput) (*TransformUrlEncodeOutput, error) {
	return &TransformUrlEncodeOutput{
		Value: strings.ReplaceAll(strings.ReplaceAll(input.Text, " ", "%20"), "&", "%26"),
	}, nil
}
