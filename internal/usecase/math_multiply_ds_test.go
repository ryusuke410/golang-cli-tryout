package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMathMultiplyInput(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name    string
		numbers []float64
		base    int
		wantErr bool
	}

	testCases := []testCase{
		{
			name:    "success with valid base",
			numbers: []float64{1.0, 2.0},
			base:    10,
			wantErr: false,
		},
		{
			name:    "success with base 0",
			numbers: []float64{1.0, 2.0},
			base:    0,
			wantErr: false,
		},
		{
			name:    "success with base 36",
			numbers: []float64{1.0, 2.0},
			base:    36,
			wantErr: false,
		},
		{
			name:    "error with invalid base (base < 0)",
			numbers: []float64{1.0, 2.0},
			base:    -1,
			wantErr: true,
		},
		{
			name:    "error with invalid base (1 < base < 2)",
			numbers: []float64{1.0, 2.0},
			base:    1,
			wantErr: true,
		},
		{
			name:    "error with invalid base (base > 36)",
			numbers: []float64{1.0, 2.0},
			base:    37,
			wantErr: true,
		},
		{
			name:    "error with one number",
			numbers: []float64{1.0},
			base:    10,
			wantErr: true,
		},
		{
			name:    "error with empty slice",
			numbers: []float64{},
			base:    10,
			wantErr: true,
		},
		{
			name:    "error with nil slice",
			numbers: nil,
			base:    10,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := NewMathMultiplyInput(tc.numbers, tc.base)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, tc.numbers, got.numbers)
				assert.Equal(t, tc.base, got.base)
			}
		})
	}
}
