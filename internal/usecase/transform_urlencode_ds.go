package usecase

type TransformUrlEncodeOutput struct {
	Value string
}

type TransformUrlEncodeInput struct {
	text string
}

func NewTransformUrlEncodeInput(text string) (*TransformUrlEncodeInput, error) {
	return &TransformUrlEncodeInput{text: text}, nil
}
