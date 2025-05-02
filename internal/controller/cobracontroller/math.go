package cobracontroller

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"ryusuke410/golang-cli-tryout/internal/usecase"
	"ryusuke410/golang-cli-tryout/pkg/fp"
)

func newMathAddCommand(mathUseCase usecase.IMathUseCase) *cobra.Command {
	c := &cobra.Command{
		Use:   "add",
		Short: "Add numbers",
		Long:  "Add multiple numbers with optional rounding",
		Args:  cobra.MinimumNArgs(2),
	}
	c.Flags().BoolP("round", "r", false, "Round the result to nearest integer")
	c.RunE = func(cmd *cobra.Command, args []string) error {
		var err error
		round, err := cmd.Flags().GetBool("round")
		if err != nil {
			return err
		}
		numbers, err := fp.MapE(args, func(arg string, _ int) (float64, error) {
			return strconv.ParseFloat(arg, 64)
		})
		if err != nil {
			return err
		}
		input, err := usecase.NewMathAddInput(numbers, round)
		if err != nil {
			return err
		}
		output, err := mathUseCase.Add(cmd.Context(), input)
		if err != nil {
			return err
		}

		fmt.Printf("Result: %.2f\n", output.Value)
		return nil
	}
	return c
}

func newMathMultiplyCommand(mathUseCase usecase.IMathUseCase) *cobra.Command {
	c := &cobra.Command{
		Use:   "multiply",
		Short: "Multiply numbers",
		Long:  "Multiply multiple numbers with optional base",
		Args:  cobra.MinimumNArgs(2),
	}
	c.Flags().IntP("base", "b", 0, "Convert result to specified base (0 for floating point)")
	c.RunE = func(cmd *cobra.Command, args []string) error {
		var err error
		base, err := cmd.Flags().GetInt("base")
		if err != nil {
			return err
		}
		numbers, err := fp.MapE(args, func(arg string, _ int) (float64, error) {
			return strconv.ParseFloat(arg, 64)
		})
		if err != nil {
			return err
		}
		input, err := usecase.NewMathMultiplyInput(numbers, base)
		if err != nil {
			return err
		}
		output, err := mathUseCase.Multiply(cmd.Context(), input)
		if err != nil {
			return err
		}

		fmt.Printf("Result: %s\n", output.Value)
		return nil
	}
	return c
}

func newMathCommand(mathUseCase usecase.IMathUseCase) *cobra.Command {
	c := &cobra.Command{
		Use:   "math",
		Short: "Perform mathematical operations",
		Long:  "Various mathematical operations with different types of arguments and options",
	}
	c.AddCommand(newMathAddCommand(mathUseCase))
	c.AddCommand(newMathMultiplyCommand(mathUseCase))
	return c
}
