package urfavecontroller

import (
	"context"
	"fmt"

	"ryusuke410/golang-cli-tryout/internal/usecase"

	"github.com/urfave/cli/v3"
)

func newTransformBase64Command(transformUseCase usecase.ITransformUseCase) *cli.Command {
	return &cli.Command{
		Name:  "base64",
		Usage: "Encode text using base64",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name: "text",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			var err error
			noNewline := cmd.Bool("no-newline")
			input, err := usecase.NewTransformBase64Input(cmd.StringArg("text"))
			if err != nil {
				return err
			}
			result, err := transformUseCase.Base64Encode(ctx, input)
			if err != nil {
				return err
			}
			fmt.Print(result.Value)
			if !noNewline {
				fmt.Println()
			}
			return nil
		},
	}
}

func newTransformUrlEncodeCommand(transformUseCase usecase.ITransformUseCase) *cli.Command {
	return &cli.Command{
		Name:  "url-encode",
		Usage: "URL encode the text",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name: "text",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			var err error
			noNewline := cmd.Bool("no-newline")
			input, err := usecase.NewTransformUrlEncodeInput(cmd.StringArg("text"))
			if err != nil {
				return err
			}
			result, err := transformUseCase.UrlEncode(ctx, input)
			if err != nil {
				return err
			}
			fmt.Print(result.Value)
			if !noNewline {
				fmt.Println()
			}
			return nil
		},
	}
}

func newTransformEscapeCommand(transformUseCase usecase.ITransformUseCase) *cli.Command {
	return &cli.Command{
		Name:  "escape",
		Usage: "Escape the text",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name: "text",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			var err error
			noNewline := cmd.Bool("no-newline")
			input, err := usecase.NewTransformEscapeInput(cmd.StringArg("text"))
			if err != nil {
				return err
			}
			result, err := transformUseCase.Escape(ctx, input)
			if err != nil {
				return err
			}
			fmt.Print(result.Value)
			if !noNewline {
				fmt.Println()
			}
			return nil
		},
	}
}

func newTransformCommand(transformUseCase usecase.ITransformUseCase) *cli.Command {
	return &cli.Command{
		Name:        "transform",
		Usage:       "Transform text using various encodings",
		Description: "Various transformations with different types of arguments and options",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "no-newline",
				Aliases: []string{"n"},
				Usage:   "Do not add newline at end of output",
				Value:   false,
			},
		},
		Commands: []*cli.Command{
			newTransformBase64Command(transformUseCase),
			newTransformUrlEncodeCommand(transformUseCase),
			newTransformEscapeCommand(transformUseCase),
		},
	}
}
