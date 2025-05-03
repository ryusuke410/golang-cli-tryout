package usecase

import (
	"context"
	"encoding/base64"
)

func (u *transformUseCase) Base64Encode(ctx context.Context, input *TransformBase64Input) (*TransformBase64Output, error) {
	return &TransformBase64Output{
		Value: base64.StdEncoding.EncodeToString([]byte(input.text)),
	}, nil
}
