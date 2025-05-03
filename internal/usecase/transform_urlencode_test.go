package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlEncode(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name    string
		text    string
		want    *TransformUrlEncodeOutput
		wantErr bool
	}

	testCases := []testCase{
		{
			name:    "success with no special characters",
			text:    "hello",
			want:    &TransformUrlEncodeOutput{Value: "hello"},
			wantErr: false,
		},
		{
			name:    "success with space",
			text:    "hello world",
			want:    &TransformUrlEncodeOutput{Value: "hello%20world"},
			wantErr: false,
		},
		{
			name:    "success with ampersand",
			text:    "key=value&key2=value2",
			want:    &TransformUrlEncodeOutput{Value: "key=value%26key2=value2"},
			wantErr: false,
		},
		{
			name:    "success with both space and ampersand",
			text:    "hello world&goodbye world",
			want:    &TransformUrlEncodeOutput{Value: "hello%20world%26goodbye%20world"},
			wantErr: false,
		},
		{
			name:    "success with empty string",
			text:    "",
			want:    &TransformUrlEncodeOutput{Value: ""},
			wantErr: false,
		},
		{
			name:    "success with multiple special characters",
			text:    "hello — world",
			want:    &TransformUrlEncodeOutput{Value: "hello%20—%20world"},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input, err := NewTransformUrlEncodeInput(tc.text)
			if err != nil {
				t.Fatalf("NewTransformUrlEncodeInput failed: %v", err)
			}

			useCase := NewTransformUseCase()
			ctx := context.Background()
			got, err := useCase.UrlEncode(ctx, input)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, tc.want.Value, got.Value)
			}
		})
	}
}
