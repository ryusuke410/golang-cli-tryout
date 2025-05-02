package usecase

// TransformBase64Output represents the output for the base64 transform operation
type TransformBase64Output struct {
	Value string
}

// TransformBase64Input represents the input for the base64 transform operation
type transformBase64Input struct {
	Text string
}

func NewTransformBase64Input(text string) (*transformBase64Input, error) {
	return &transformBase64Input{Text: text}, nil
}
