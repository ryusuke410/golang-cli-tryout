package service

import (
	"fmt"
	"strconv"

	"ryusuke410/golang-cli-tryout/internal/domain/service"
)

// mathService implements IMathService
type mathService struct{}

var _ service.IMathService = (*mathService)(nil)

// NewMathService creates a new math service
func NewMathService() *mathService {
	return &mathService{}
}

func (s *mathService) Sum(numbers []float64) float64 {
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func (s *mathService) Product(numbers []float64) float64 {
	if len(numbers) < 2 {
		panic("at least 2 numbers are required")
	}
	var product float64 = 1
	for _, num := range numbers {
		product *= num
	}
	return product
}

func (s *mathService) Round(num float64) float64 {
	return float64(int(num + 0.5))
}

func (s *mathService) FormatNumber(num float64, base int) string {
	if base <= 0 {
		return fmt.Sprintf("%.2f", num)
	}
	return strconv.FormatInt(int64(num), base)
}
