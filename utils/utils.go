package utils

import (
	"fmt"

	"gonum.org/v1/gonum/optimize"
)

// Config содержит параметры для оптимизации
type Config struct {
	// MaxIterations максимальное количество итераций
	MaxIterations int
	// GradientThreshold порог для градиента
	GradientThreshold float64
}

// Конфигурация по умолчанию
func DefaultConfig() Config {
	return Config{
		MaxIterations:     1000,
		GradientThreshold: 1e-8,
	}
}

// Validate проверяет корректность конфигурации
func (cfg Config) Validate() error {
	if cfg.MaxIterations <= 0 {
		return fmt.Errorf("max iterations must be positive: %d", cfg.MaxIterations)
	}
	if cfg.GradientThreshold <= 0 {
		return fmt.Errorf("gradient threshold must be positive: %e", cfg.GradientThreshold)
	}
	return nil
}

// Решение системы нелинейных уравнений
func Root(equations []func([]float64) float64, x0 []float64, config ...Config) (*optimize.Result, error) {
	// Валидация входных параметров
	if len(equations) == 0 {
		return nil, fmt.Errorf("system non linear")
	}
	if len(x0) == 0 {
		return nil, fmt.Errorf("initial guess cannot be empty")
	}

	problem := optimize.Problem{
		Func: func(x []float64) float64 {
			var sumSq float64
			for _, f := range equations {
				res := f(x)
				sumSq += res * res
			}
			return sumSq
		},
	}

	// Используем конфигурацию по умолчанию, если не передана
	cfg := DefaultConfig()
	if len(config) > 0 {
		cfg = config[0]
		if err := cfg.Validate(); err != nil {
			return nil, fmt.Errorf("system non linear")
		}
	}

	solver := &optimize.NelderMead{} // не требует градиента

	settings := &optimize.Settings{
		MajorIterations:   cfg.MaxIterations,
		GradientThreshold: cfg.GradientThreshold,
	}

	result, err := optimize.Minimize(problem, x0, settings, solver)
	if err != nil {
		return &optimize.Result{}, fmt.Errorf("optimization failed: %w", err)
	}

	return result, nil
}
