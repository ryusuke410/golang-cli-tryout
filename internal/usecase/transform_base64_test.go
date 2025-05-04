package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64Encode(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name    string
		text    string
		want    *TransformBase64Output
		wantErr bool
	}

	testCases := []testCase{
		{
			name:    "success with valid text",
			text:    "hello",
			want:    &TransformBase64Output{Value: "aGVsbG8="},
			wantErr: false,
		},
		{
			name:    "success with empty text",
			text:    "",
			want:    &TransformBase64Output{Value: ""},
			wantErr: false,
		},
		{
			name:    "success with special characters",
			text:    "hello!@#$%",
			want:    &TransformBase64Output{Value: "aGVsbG8hQCMkJQ=="},
			wantErr: false,
		},
		{
			name:    "success with unicode",
			text:    "こんにちは",
			want:    &TransformBase64Output{Value: "44GT44KT44Gr44Gh44Gv"},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input, err := NewTransformBase64Input(tc.text)
			if err != nil {
				t.Fatalf("NewTransformBase64Input failed: %v", err)
			}

			u := NewTransformUseCase()
			got, err := u.Base64Encode(context.Background(), input)

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
