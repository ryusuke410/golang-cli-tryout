package service

import (
	"fmt"
	"math"
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

func (s *mathService) Product(numbers []float64) (float64, error) {
	if len(numbers) < 2 {
		return 0, fmt.Errorf("at least 2 numbers are required")
	}
	var product float64 = 1
	for _, num := range numbers {
		product *= num
	}
	return product, nil
}

func (s *mathService) Round(num float64) float64 {
	return math.Round(num)
}

func (s *mathService) FormatNumber(num float64, base int) (string, error) {
	if !(base == 0 || (2 <= base && base <= 36)) {
		return "", fmt.Errorf("invalid base")
	}
	if base == 0 {
		return fmt.Sprintf("%.2f", num), nil
	}
	return strconv.FormatInt(int64(num+0.5), base), nil
}
