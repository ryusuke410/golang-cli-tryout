package usecase

import (
	"context"
	"strings"
)

func (u *transformUseCase) Escape(ctx context.Context, input *TransformEscapeInput) (*TransformEscapeOutput, error) {
	return &TransformEscapeOutput{
		Value: strings.ReplaceAll(strings.ReplaceAll(input.text, `\`, `\\`), `"`, `\"`),
	}, nil
}
