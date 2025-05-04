package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTransformUrlEncodeInput(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		wantErr bool
	}{
		{
			name:    "valid input",
			text:    "hello world",
			wantErr: false,
		},
		{
			name:    "empty input",
			text:    "",
			wantErr: false,
		},
		{
			name:    "special characters",
			text:    "hello!@#$%^&*()",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTransformUrlEncodeInput(tt.text)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, tt.text, got.text)
			}
		})
	}
}
