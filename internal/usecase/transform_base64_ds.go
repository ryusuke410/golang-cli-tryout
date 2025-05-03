package usecase

type TransformBase64Output struct {
	Value string
}

type TransformBase64Input struct {
	text string
}

func NewTransformBase64Input(text string) (*TransformBase64Input, error) {
	return &TransformBase64Input{text: text}, nil
}
