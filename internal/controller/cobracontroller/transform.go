package cobracontroller

import (
	"fmt"

	"github.com/spf13/cobra"

	"ryusuke410/golang-cli-tryout/internal/usecase"
)

func newTransformBase64Command(transformUseCase usecase.ITransformUseCase) *cobra.Command {
	c := &cobra.Command{
		Use:   "base64",
		Short: "Encode text using base64",
		Long:  "Encode text using base64 encoding",
		Args:  cobra.ExactArgs(1),
	}
	c.RunE = func(cmd *cobra.Command, args []string) error {
		var err error
		text := args[0]
		noNewline, err := cmd.Flags().GetBool("no-newline")
		if err != nil {
			return err
		}
		input, err := usecase.NewTransformBase64Input(text)
		if err != nil {
			return err
		}
		result, err := transformUseCase.Base64Encode(cmd.Context(), input)
		if err != nil {
			return err
		}
		printResult(result.Value, noNewline)
		return nil
	}
	return c
}

func newTransformUrlEncodeCommand(transformUseCase usecase.ITransformUseCase) *cobra.Command {
	c := &cobra.Command{
		Use:   "url-encode",
		Short: "URL encode the text",
		Long:  "URL encode the text",
		Args:  cobra.ExactArgs(1),
	}
	c.RunE = func(cmd *cobra.Command, args []string) error {
		var err error
		text := args[0]
		noNewline, err := cmd.Flags().GetBool("no-newline")
		if err != nil {
			return err
		}
		input, err := usecase.NewTransformUrlEncodeInput(text)
		if err != nil {
			return err
		}
		result, err := transformUseCase.UrlEncode(cmd.Context(), input)
		if err != nil {
			return err
		}
		printResult(result.Value, noNewline)
		return nil
	}
	return c
}

func newTransformEscapeCommand(transformUseCase usecase.ITransformUseCase) *cobra.Command {
	c := &cobra.Command{
		Use:   "escape",
		Short: "Escape special characters",
		Long:  "Escape special characters",
		Args:  cobra.ExactArgs(1),
	}
	c.RunE = func(cmd *cobra.Command, args []string) error {
		var err error
		text := args[0]
		noNewline, err := cmd.Flags().GetBool("no-newline")
		if err != nil {
			return err
		}
		input, err := usecase.NewTransformEscapeInput(text)
		if err != nil {
			return err
		}
		result, err := transformUseCase.Escape(cmd.Context(), input)
		if err != nil {
			return err
		}
		printResult(result.Value, noNewline)
		return nil
	}
	return c
}

func newTransformCommand(transformUseCase usecase.ITransformUseCase) *cobra.Command {
	c := &cobra.Command{
		Use:   "transform",
		Short: "Transform text using various encodings",
		Long: `Transform text using various encodings and transformations.
Available transformations: base64, url-encode, escape`,
	}
	c.PersistentFlags().BoolP("no-newline", "n", false, "Do not add newline at end of output")
	c.AddCommand(newTransformBase64Command(transformUseCase))
	c.AddCommand(newTransformUrlEncodeCommand(transformUseCase))
	c.AddCommand(newTransformEscapeCommand(transformUseCase))
	return c
}

func printResult(value string, noNewline bool) {
	fmt.Print(value)
	if !noNewline {
		fmt.Println()
	}
}
