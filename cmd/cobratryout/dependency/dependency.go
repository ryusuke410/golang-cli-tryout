package dependency

import (
	"github.com/spf13/cobra"

	"ryusuke410/golang-cli-tryout/internal/controller/cobracontroller"
	"ryusuke410/golang-cli-tryout/internal/infrastructure/service"
	"ryusuke410/golang-cli-tryout/internal/usecase"
)

func New() *cobra.Command {
	mathService := service.NewMathService()
	mathUseCase := usecase.NewMathUseCase(mathService)
	transformUseCase := usecase.NewTransformUseCase()

	return cobracontroller.NewRootCommand(mathUseCase, transformUseCase)
}
