package urfavecontroller

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"ryusuke410/golang-cli-tryout/internal/usecase"
)

func newMathAddCommand(mathUseCase usecase.IMathUseCase) *cli.Command {
	return &cli.Command{
		Name:        "add",
		Usage:       "Add numbers",
		Description: "Add multiple numbers with optional rounding",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "round",
				Aliases: []string{"r"},
				Usage:   "Round the result to nearest integer",
				Value:   false,
			},
		},
		Arguments: []cli.Argument{
			&cli.FloatArgs{
				Name: "numbers",
				Min:  2,  // 最低2つの引数が必要
				Max:  -1, // 上限なし
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			input, err := usecase.NewMathAddInput(cmd.FloatArgs("numbers"), cmd.Bool("round"))
			if err != nil {
				return err
			}

			output, err := mathUseCase.Add(ctx, input)
			if err != nil {
				return err
			}

			fmt.Printf("Result: %.2f\n", output.Value)
			return nil
		},
	}
}

func newMathMultiplyCommand(mathUseCase usecase.IMathUseCase) *cli.Command {
	return &cli.Command{
		Name:        "multiply",
		Usage:       "Multiply numbers",
		Description: "Multiply multiple numbers with optional base",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "base",
				Aliases: []string{"b"},
				Usage:   "Convert result to specified base (0 for floating point)",
				Value:   0,
			},
		},
		Arguments: []cli.Argument{
			&cli.FloatArgs{
				Name: "numbers",
				Min:  2,  // 最低2つの引数が必要
				Max:  -1, // 上限なし
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			input, err := usecase.NewMathMultiplyInput(cmd.FloatArgs("numbers"), cmd.Int("base"))
			if err != nil {
				return err
			}

			output, err := mathUseCase.Multiply(ctx, input)
			if err != nil {
				return err
			}

			fmt.Printf("Result: %s\n", output.Value)
			return nil
		},
	}
}

func newMathCommand(mathUseCase usecase.IMathUseCase) *cli.Command {
	return &cli.Command{
		Name:        "math",
		Usage:       "Perform mathematical operations",
		Description: "Various mathematical operations with different types of arguments and options",
		Commands: []*cli.Command{
			newMathAddCommand(mathUseCase),
			newMathMultiplyCommand(mathUseCase),
		},
	}
}
