package urfavecontroller_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"ryusuke410/golang-cli-tryout/internal/controller/urfavecontroller"
	"ryusuke410/golang-cli-tryout/internal/mocks/mockusecase"
)

func TestRootCommand(t *testing.T) {
	var err error

	mockMathUseCase := mockusecase.NewIMathUseCase(t)
	mockTransformUseCase := mockusecase.NewITransformUseCase(t)

	rootCmd := urfavecontroller.NewRootCommand(mockMathUseCase, mockTransformUseCase)

	ctx := context.Background()
	args := []string{"urfavetryout"}

	err = rootCmd.Run(ctx, args)
	assert.NoError(t, err)
}
