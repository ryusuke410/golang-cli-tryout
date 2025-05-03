package usecase

import (
	"context"
	"testing"

	"ryusuke410/golang-cli-tryout/internal/infrastructure/service"

	"github.com/stretchr/testify/assert"
)

func TestMathMultiply(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name    string
		numbers []float64
		base    int
		want    *MathMultiplyOutput
		wantErr bool
	}

	testCases := []testCase{
		{
			name:    "success with valid input",
			numbers: []float64{2.0, 3.0},
			base:    10,
			want:    &MathMultiplyOutput{Value: "6"},
			wantErr: false,
		},
		{
			name:    "success with zero",
			numbers: []float64{0.0, 5.0},
			base:    10,
			want:    &MathMultiplyOutput{Value: "0"},
			wantErr: false,
		},
		{
			name:    "success with different base",
			numbers: []float64{2.0, 3.0},
			base:    16,
			want:    &MathMultiplyOutput{Value: "6"},
			wantErr: false,
		},
		{
			name:    "success with multiple numbers",
			numbers: []float64{2.0, 3.0, 4.0},
			base:    10,
			want:    &MathMultiplyOutput{Value: "24"},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input, err := NewMathMultiplyInput(tc.numbers, tc.base)
			if err != nil {
				t.Fatalf("NewMathMultiplyInput failed: %v", err)
			}

			// Create real service
			mathService := service.NewMathService()

			// Create use case with real service
			u := NewMathUseCase(mathService)
			got, err := u.Multiply(context.Background(), input)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}
