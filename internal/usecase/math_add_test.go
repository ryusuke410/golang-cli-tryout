package usecase

import (
	"context"
	"testing"

	"ryusuke410/golang-cli-tryout/internal/infrastructure/service"

	"github.com/stretchr/testify/assert"
)

func TestMathAdd(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name    string
		numbers []float64
		round   bool
		want    *MathAddOutput
		wantErr bool
	}

	testCases := []testCase{
		{
			name:    "success with normal numbers and round=false",
			numbers: []float64{1.0, 2.0},
			round:   false,
			want: &MathAddOutput{
				Value: 3.0,
			},
			wantErr: false,
		},
		{
			name:    "success with normal numbers and round=true",
			numbers: []float64{1.0, 2.5},
			round:   true,
			want: &MathAddOutput{
				Value: 4.0,
			},
			wantErr: false,
		},
		{
			name:    "success with multiple numbers",
			numbers: []float64{1.0, 2.0, 3.0},
			round:   false,
			want: &MathAddOutput{
				Value: 6.0,
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create real service
			mathService := service.NewMathService()

			// Create use case with real service
			u := NewMathUseCase(mathService)

			// Create input using NewMathAddInput
			input, err := NewMathAddInput(tc.numbers, tc.round)
			if err != nil {
				t.Fatalf("NewMathAddInput failed: %v", err)
			}

			// Execute test
			got, err := u.Add(context.Background(), input)

			// Verify results
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
