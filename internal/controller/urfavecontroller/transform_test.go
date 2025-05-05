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

func TestTransformBase64Command(t *testing.T) {
	var err error

	mockTransformUseCase := mockusecase.NewITransformUseCase(t)
	mockMathUseCase := mockusecase.NewIMathUseCase(t)

	mockTransformUseCase.EXPECT().Base64Encode(mock.Anything, mock.Anything).Return(&usecase.TransformBase64Output{Value: "dGVzdA=="}, nil)

	rootCmd := urfavecontroller.NewRootCommand(mockMathUseCase, mockTransformUseCase)

	ctx := context.Background()
	args := []string{"urfavetryout", "transform", "base64", "test"}

	err = rootCmd.Run(ctx, args)
	assert.NoError(t, err)
}

func TestTransformUrlEncodeCommand(t *testing.T) {
	var err error

	mockTransformUseCase := mockusecase.NewITransformUseCase(t)
	mockMathUseCase := mockusecase.NewIMathUseCase(t)

	mockTransformUseCase.EXPECT().UrlEncode(mock.Anything, mock.Anything).Return(&usecase.TransformUrlEncodeOutput{Value: "test%20data"}, nil)

	rootCmd := urfavecontroller.NewRootCommand(mockMathUseCase, mockTransformUseCase)

	ctx := context.Background()
	args := []string{"urfavetryout", "transform", "url-encode", "test data"}

	err = rootCmd.Run(ctx, args)
	assert.NoError(t, err)
}
