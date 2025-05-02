package urfavecontroller

import (
	"github.com/urfave/cli/v3"

	"ryusuke410/golang-cli-tryout/internal/usecase"
)

func NewRootCommand(mathUseCase usecase.IMathUseCase, transformUseCase usecase.ITransformUseCase) *cli.Command {
	return &cli.Command{
		Name:        "urfavetryout",
		Usage:       "A CLI tool for math and transform operations",
		Description: "A CLI tool for math and transform operations",
		Commands: []*cli.Command{
			newMathCommand(mathUseCase),
			newTransformCommand(transformUseCase),
		},
	}
}
