package cvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseValidator(t *testing.T) {
	t.Parallel()

	type input struct {
		base int
	}

	testCases := []struct {
		name     string
		input    input
		wantPass bool
	}{
		{
			name: "success with base 0",
			input: input{
				base: 0,
			},
			wantPass: true,
		},
		{
			name: "success with base 2",
			input: input{
				base: 2,
			},
			wantPass: true,
		},
		{
			name: "success with base 36",
			input: input{
				base: 36,
			},
			wantPass: true,
		},
		{
			name: "error with base 1",
			input: input{
				base: 1,
			},
			wantPass: false,
		},
		{
			name: "error with base 37",
			input: input{
				base: 37,
			},
			wantPass: false,
		},
		{
			name: "error with negative base",
			input: input{
				base: -1,
			},
			wantPass: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			testStruct := struct {
				base int `validate:"c_base"`
			}{
				base: tc.input.base,
			}

			err := Struct(testStruct)
			if tc.wantPass {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
