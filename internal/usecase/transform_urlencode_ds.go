package usecase

// TransformUrlEncodeOutput represents the output for the url encode transform operation
type TransformUrlEncodeOutput struct {
	Value string
}

// TransformUrlEncodeInput represents the input for the url encode transform operation
type transformUrlEncodeInput struct {
	Text string
}

func NewTransformUrlEncodeInput(text string) (*transformUrlEncodeInput, error) {
	return &transformUrlEncodeInput{Text: text}, nil
}
