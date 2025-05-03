package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTransformBase64Input(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name    string
		text    string
		wantErr bool
	}

	testCases := []testCase{
		{
			name:    "success with valid text",
			text:    "hello",
			wantErr: false,
		},
		{
			name:    "success with empty text",
			text:    "",
			wantErr: false,
		},
		{
			name:    "success with special characters",
			text:    "hello!@#$%",
			wantErr: false,
		},
		{
			name:    "success with unicode",
			text:    "こんにちは",
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := NewTransformBase64Input(tc.text)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, tc.text, got.text)
			}
		})
	}
}
