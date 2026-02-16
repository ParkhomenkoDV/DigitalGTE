package system

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/diff/fd"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/optimize"
)

// Система уравнений
type System struct {
	Equations []func([]float64) float64
}

// Создание нрвой системы уравнений
func New(equations []func([]float64) float64) (*System, error) {
	if len(equations) == 0 {
		return &System{}, fmt.Errorf("empty system equations")
	}

	// Проверяем, что все функции не nil
	for i, eq := range equations {
		if eq == nil {
			return nil, fmt.Errorf("equation at index %d is nil", i)
		}
	}

	return &System{
		Equations: equations,
	}, nil
}

// Вычисление вектора невязок
func (s *System) F(x []float64) []float64 {
	res := make([]float64, len(s.Equations))
	for i, eq := range s.Equations {
		res[i] = eq(x)
	}
	return res
}

// Критерий сходимости
func (s *System) Func(x []float64) float64 {
	var sumSq float64
	for _, eq := range s.Equations {
		f := eq(x)
		sumSq += f * f
	}
	return sumSq
}

// Численное вычисление матрицы Якоби
func (s *System) Jacobian(x []float64) *mat.Dense {
	m := len(s.Equations) // количество уравнений
	n := len(x)           // количество переменных

	jac := mat.NewDense(m, n, nil)

	// Функция для численного дифференцирования
	f := func(y, x []float64) {
		for i := range s.Equations {
			y[i] = s.Equations[i](x)
		}
	}

	// Вычисляем якобиан методом центральных разностей
	fd.Jacobian(jac, f, x, &fd.JacobianSettings{
		Formula:    fd.Central, // Центральные разности (наиболее точные)
		Step:       1e-8,       // Шаг дифференцирования
		Concurrent: true,       // Параллельное вычисление
	})

	return jac
}

// Численное вычисление градиента
func (s *System) Grad(grad, x []float64) {
	// Используем численное дифференцирование
	fd.Gradient(grad, s.Func, x, &fd.Settings{
		Formula:    fd.Central, // Центральные разности (наиболее точные)
		Step:       1e-8,       // Шаг дифференцирования
		Concurrent: true,       // Параллельное вычисление
	})
}

// Численное вычисление градиента
func (s *System) Hess(hess *mat.SymDense, x []float64) {
	// Используем численное дифференцирование
	fd.Hessian(hess, s.Func, x, &fd.Settings{
		Formula:    fd.Central, // Центральные разности (наиболее точные)
		Step:       1e-8,       // Шаг дифференцирования
		Concurrent: true,       // Параллельное вычисление
	})
}

var (
	// NoGrad solvers
	NelderMead optimize.Method = &optimize.NelderMead{}              // медленный и неточный
	CmaEsChol  optimize.Method = &optimize.CmaEsChol{Population: 20} // очень неточный
	// Grad solvers
	LBFGS           optimize.Method = &optimize.LBFGS{}
	BFGS            optimize.Method = &optimize.BFGS{}
	GradientDescent optimize.Method = &optimize.GradientDescent{GradStopThreshold: 0.000_000_1} // неточен, застревает в локальных минимумах
	CG              optimize.Method = &optimize.CG{}
	// Grad + Hess solvers
	Newton optimize.Method = &optimize.Newton{}
)

// Solve решает систему методом Ньютона
func (s *System) Solve(x0 []float64, method optimize.Method) (*optimize.Result, error) {
	if len(x0) == 0 {
		return &optimize.Result{}, fmt.Errorf("empty initial guess")
	}

	problem := optimize.Problem{
		Func: s.Func,
		Grad: s.Grad,
		Hess: s.Hess,
	}

	// Настройки оптимизации
	settings := &optimize.Settings{
		MajorIterations:   1_000, // Максимальное число итераций
		GradientThreshold: 1e-8,  // Порог для градиента
		Concurrent:        0,     // Автоматический выбор числа потоков
	}

	return optimize.Minimize(problem, x0, settings, method)
}

// Проверка сходимости решения
func (s *System) IsConvergence(x []float64, tolerance float64) bool {
	residuals := s.F(x)
	for _, r := range residuals {
		if math.Abs(r) > tolerance {
			return false
		}
	}

	return true
}
