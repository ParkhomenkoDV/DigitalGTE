package system

import (
	"math"
	"strconv"
	"testing"

	"gonum.org/v1/gonum/optimize"
)

var tests = []struct {
	name      string
	equations []func([]float64) float64
	x0        []float64
	want      []float64
}{
	{
		name: "1",
		equations: []func([]float64) float64{
			func(x []float64) float64 {
				return x[0]*x[0] - 2
			},
		},
		x0:   []float64{1.0},
		want: []float64{math.Sqrt(2)},
	},
	{
		name: "2",
		equations: []func([]float64) float64{
			func(x []float64) float64 {
				return x[0]*x[0] + x[1]*x[1] - 4
			},
			func(x []float64) float64 {
				return x[0] - x[1]
			},
		},
		x0:   []float64{1.0, 1.0},
		want: []float64{math.Sqrt(2), math.Sqrt(2)},
	},
	{
		name: "3",
		equations: []func([]float64) float64{
			func(x []float64) float64 {
				return x[0] + x[1] + x[2] - 6
			},
			func(x []float64) float64 {
				return x[0]*x[0] + x[1]*x[1] + x[2]*x[2] - 14
			},
			func(x []float64) float64 {
				return x[0]*x[1]*x[2] - 6
			},
		},
		x0:   []float64{0.5, 1.5, 2.5},
		want: []float64{1.0, 2.0, 3.0},
	},
	{
		name: "tri",
		equations: []func([]float64) float64{
			func(x []float64) float64 {
				return math.Sin(x[0]) - 0.5
			},
		},
		x0:   []float64{0.5},
		want: []float64{math.Pi / 6},
	},
	{
		name: "exp",
		equations: []func([]float64) float64{
			func(x []float64) float64 {
				return math.Exp(x[0]) - 2
			},
		},
		x0:   []float64{0.5},
		want: []float64{math.Log(2)},
	},
	{
		name:      "empty equations",
		equations: []func([]float64) float64{},
		x0:        []float64{1.0},
	},
	{
		name: "empty initial guess",
		equations: []func([]float64) float64{
			func(x []float64) float64 { return x[0]*x[0] - 2 },
		},
		x0: []float64{},
	},
}

func TestF(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s, err := New(test.equations)

			if err != nil && len(test.equations) == 0 {
				return // ошибка предполагалась
			}

			if len(test.x0) == 0 {
				return // ошибка предполагалась
			}

			want := make([]float64, len(s.Equations))
			for i := 0; i < len(s.Equations); i++ {
				want[i] = s.Equations[i](test.x0)
			}
			got := s.F(test.x0)
			for i := range want {
				if got[i] != want[i] {
					t.Errorf("F() = %v, want: %v", got, want)
					return
				}
			}
		})
	}
}

var methods = []optimize.Method{
	// NelderMead, // медленный и неточный
	// CmaEsChol,  // очень неточный
	LBFGS,
	BFGS,
	// GradientDescent, // неточен, застревает в локальных минимумах
	CG,
	Newton,
}

// Тест решения системы нелинейных уравнений
func TestSolve(t *testing.T) {
	for _, test := range tests {
		for _, method := range methods {
			t.Run(test.name, func(t *testing.T) {
				s, err := New(test.equations)

				if err != nil && len(test.equations) == 0 {
					return // ошибка предполагалась
				}

				result, err := s.Solve(test.x0, method)

				if err != nil && len(test.x0) == 0 {
					return // ошибка предполагалась
				}

				if err != nil {
					t.Errorf("Solve(%v) error = %v", method, err)
					return
				}

				if result == nil {
					t.Error("Solve() returned nil result")
					return
				}

				if !s.IsConvergence(result.X, 1e-6) {
					t.Errorf("Solve().X = %f, want %f", result.X, test.want)
				}
			})
		}
	}
}

// Бенчмарк решения системы нелинейных уравнений
func BenchmarkSolve(b *testing.B) {
	for _, test := range tests {
		if len(test.equations) == 0 || len(test.x0) == 0 {
			continue
		}

		s, err := New(test.equations)
		if err != nil {
			b.Skipf("Skipping %s: %v", test.name, err)
		}

		for j, method := range methods {
			b.Run(test.name+strconv.Itoa(j), func(b *testing.B) {

				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					// Создаем копию начального приближения для каждого запуска
					x0Copy := make([]float64, len(test.x0))
					copy(x0Copy, test.x0)

					_, err = s.Solve(x0Copy, method)
					if err != nil {
						b.Errorf("Solve() error = %v", err)
					}
				}
			})
		}
	}
}
