package cvalidator

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestRegisterValidation(t *testing.T) {
	t.Parallel()
	RegisterValidation("test_base", func(fl validator.FieldLevel) bool {
		base := fl.Field().Int()
		return base == 0 || (2 <= base && base <= 36)
	})

	type input struct {
		Numbers []float64 `validate:"min=2"`
		Base    int       `validate:"test_base"`
	}

	testCases := []struct {
		name    string
		input   input
		wantErr bool
	}{
		{
			name: "success with valid numbers and base",
			input: input{
				Numbers: []float64{1.0, 2.0},
				Base:    10,
			},
		},
		{
			name: "error with invalid numbers",
			input: input{
				Numbers: []float64{1.0},
				Base:    10,
			},
			wantErr: true,
		},
		{
			name: "error with invalid base",
			input: input{
				Numbers: []float64{1.0, 2.0},
				Base:    37,
			},
			wantErr: true,
		},
		{
			name: "success with base 0",
			input: input{
				Numbers: []float64{1.0, 2.0},
				Base:    0,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := Struct(tc.input)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestStruct(t *testing.T) {
	t.Parallel()
	RegisterValidation("test_base", func(fl validator.FieldLevel) bool {
		base := fl.Field().Int()
		return base == 0 || (2 <= base && base <= 36)
	})

	type input struct {
		Numbers []float64 `validate:"min=2"`
		Base    int       `validate:"test_base"`
	}

	testCases := []struct {
		name    string
		input   input
		wantErr bool
	}{
		{
			name: "success with valid numbers and base",
			input: input{
				Numbers: []float64{1.0, 2.0},
				Base:    10,
			},
		},
		{
			name: "error with invalid numbers",
			input: input{
				Numbers: []float64{1.0},
				Base:    10,
			},
			wantErr: true,
		},
		{
			name: "error with invalid base",
			input: input{
				Numbers: []float64{1.0, 2.0},
				Base:    37,
			},
			wantErr: true,
		},
		{
			name: "success with base 0",
			input: input{
				Numbers: []float64{1.0, 2.0},
				Base:    0,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := Struct(tc.input)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
