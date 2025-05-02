package service

// IMathService represents the math service interface
type IMathService interface {
	Sum(numbers []float64) float64
	Product(numbers []float64) (float64, error)
	Round(num float64) float64
	FormatNumber(num float64, base int) string
}
