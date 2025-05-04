package urfavecontroller_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ryusuke410/golang-cli-tryout/internal/controller/urfavecontroller"
	"ryusuke410/golang-cli-tryout/internal/mocks/mockusecase"
	"ryusuke410/golang-cli-tryout/internal/usecase"
)

func TestMathAddCommand(t *testing.T) {
	var err error

	mockMathUseCase := mockusecase.NewIMathUseCase(t)
	mockTransformUseCase := mockusecase.NewITransformUseCase(t)

	mockMathUseCase.EXPECT().Add(mock.Anything, mock.Anything).Return(&usecase.MathAddOutput{Value: 5.0}, nil)

	rootCmd := urfavecontroller.NewRootCommand(mockMathUseCase, mockTransformUseCase)

	ctx := context.Background()
	args := []string{"urfavetryout", "math", "add", "2", "3"}

	err = rootCmd.Run(ctx, args)
	assert.NoError(t, err)
}

func TestMathMultiplyCommand(t *testing.T) {
	var err error

	mockMathUseCase := mockusecase.NewIMathUseCase(t)
	mockTransformUseCase := mockusecase.NewITransformUseCase(t)

	mockMathUseCase.EXPECT().Multiply(mock.Anything, mock.Anything).Return(&usecase.MathMultiplyOutput{Value: "6.0"}, nil)

	rootCmd := urfavecontroller.NewRootCommand(mockMathUseCase, mockTransformUseCase)

	ctx := context.Background()
	args := []string{"urfavetryout", "math", "multiply", "2", "3"}

	err = rootCmd.Run(ctx, args)
	assert.NoError(t, err)
}
