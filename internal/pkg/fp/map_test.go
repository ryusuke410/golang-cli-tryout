package fp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapE(t *testing.T) {
	t.Parallel()

	type input struct {
		input  []int
		mapper func(int, int) (int, error)
	}

	testCases := []struct {
		name    string
		input   input
		want    []int
		wantErr bool
	}{
		{
			name: "success with simple mapping",
			input: input{
				input: []int{1, 2, 3},
				mapper: func(v int, _ int) (int, error) {
					return v * 2, nil
				},
			},
			want: []int{2, 4, 6},
		},
		{
			name: "success with empty input",
			input: input{
				input: []int{},
				mapper: func(v int, _ int) (int, error) {
					return v, nil
				},
			},
			want: []int{},
		},
		{
			name: "success with index",
			input: input{
				input: []int{1, 2, 3},
				mapper: func(v int, i int) (int, error) {
					return v + i, nil
				},
			},
			want: []int{1, 3, 5},
		},
		{
			name: "error case",
			input: input{
				input: []int{1, 2, 3},
				mapper: func(v int, _ int) (int, error) {
					if v == 2 {
						return 0, fmt.Errorf("error on value 2")
					}
					return v, nil
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := MapE(tc.input.input, tc.input.mapper)

			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
