package usecase

import (
	"context"
	"strings"
)

func (u *transformUseCase) Escape(ctx context.Context, input *transformEscapeInput) (*TransformEscapeOutput, error) {
	return &TransformEscapeOutput{
		Value: strings.ReplaceAll(strings.ReplaceAll(input.Text, `\`, `\\`), `"`, `\"`),
	}, nil
}
