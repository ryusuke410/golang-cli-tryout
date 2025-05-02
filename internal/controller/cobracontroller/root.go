package cobracontroller

import (
	"fmt"

	"github.com/spf13/cobra"

	"ryusuke410/golang-cli-tryout/internal/usecase"
)

func NewRootCommand(mathUseCase usecase.IMathUseCase, transformUseCase usecase.ITransformUseCase) *cobra.Command {
	c := &cobra.Command{
		Use:   "cobra-tryout",
		Short: "A CLI tool for trying out various cobra features",
		Long: `A CLI tool that demonstrates various cobra features including:
- Subcommands
- Variable arguments
- Flags and options
- Short and long options`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to cobra-tryout!")
			fmt.Println("Use 'cobra-tryout help' for more information about available commands.")
		},
	}
	c.AddCommand(newMathCommand(mathUseCase))
	c.AddCommand(newTransformCommand(transformUseCase))
	return c
}
