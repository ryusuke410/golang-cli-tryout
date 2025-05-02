package dependency

import (
	"github.com/urfave/cli/v3"

	"ryusuke410/golang-cli-tryout/internal/controller/urfavecontroller"
	"ryusuke410/golang-cli-tryout/internal/infrastructure/service"
	"ryusuke410/golang-cli-tryout/internal/usecase"
)

func New() *cli.Command {
	mathService := service.NewMathService()
	mathUseCase := usecase.NewMathUseCase(mathService)
	transformUseCase := usecase.NewTransformUseCase()
	return urfavecontroller.NewRootCommand(mathUseCase, transformUseCase)
}
