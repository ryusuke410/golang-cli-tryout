package usecase

// TransformEscapeOutput represents the output for the escape transform operation
type TransformEscapeOutput struct {
	Value string
}

// TransformEscapeInput represents the input for the escape transform operation
type transformEscapeInput struct {
	Text string
}

func NewTransformEscapeInput(text string) (*transformEscapeInput, error) {
	return &transformEscapeInput{Text: text}, nil
}
