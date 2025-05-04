package usecase

type TransformEscapeOutput struct {
	Value string
}

type TransformEscapeInput struct {
	text string
}

func NewTransformEscapeInput(text string) (*TransformEscapeInput, error) {
	return &TransformEscapeInput{text: text}, nil
}
