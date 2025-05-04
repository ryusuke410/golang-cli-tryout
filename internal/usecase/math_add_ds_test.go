package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMathAddInput(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name    string
		numbers []float64
		round   bool
		wantErr bool
	}

	testCases := []testCase{
		{
			name:    "success with two numbers",
			numbers: []float64{1.0, 2.0},
			round:   true,
			wantErr: false,
		},
		{
			name:    "success with three numbers",
			numbers: []float64{1.0, 2.0, 3.0},
			round:   false,
			wantErr: false,
		},
		{
			name:    "error with one number",
			numbers: []float64{1.0},
			round:   true,
			wantErr: true,
		},
		{
			name:    "error with empty slice",
			numbers: []float64{},
			round:   true,
			wantErr: true,
		},
		{
			name:    "error with nil slice",
			numbers: nil,
			round:   true,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := NewMathAddInput(tc.numbers, tc.round)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, tc.numbers, got.numbers)
				assert.Equal(t, tc.round, got.round)
			}
		})
	}
}
