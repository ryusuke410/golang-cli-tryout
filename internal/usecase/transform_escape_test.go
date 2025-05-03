package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEscape(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name    string
		text    string
		want    *TransformEscapeOutput
		wantErr bool
	}

	testCases := []testCase{
		{
			name:    "success with no special characters",
			text:    "hello",
			want:    &TransformEscapeOutput{Value: "hello"},
			wantErr: false,
		},
		{
			name:    "success with backslash",
			text:    "hello\\world",
			want:    &TransformEscapeOutput{Value: "hello\\\\world"},
			wantErr: false,
		},
		{
			name:    "success with double quote",
			text:    "hello\"world\\",
			want:    &TransformEscapeOutput{Value: "hello\\\"world\\\\"},
			wantErr: false,
		},
		{
			name:    "success with both special characters",
			text:    "hello\"world\\",
			want:    &TransformEscapeOutput{Value: "hello\\\"world\\\\"},
			wantErr: false,
		},
		{
			name:    "success with multiple special characters",
			text:    "hello\"\"world\\\\",
			want:    &TransformEscapeOutput{Value: "hello\\\"\\\"world\\\\\\\\"},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input, err := NewTransformEscapeInput(tc.text)
			if err != nil {
				t.Fatalf("NewTransformEscapeInput failed: %v", err)
			}

			u := NewTransformUseCase()
			got, err := u.Escape(context.Background(), input)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}
