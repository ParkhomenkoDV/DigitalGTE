package utils

import (
	"math"
	"testing"
)

var tests = []struct {
	name      string
	equations []func([]float64) float64
	x0        []float64
	wantErr   bool
	want      []float64
}{
	{
		name: "1",
		equations: []func([]float64) float64{
			func(x []float64) float64 {
				return x[0]*x[0] - 2
			},
		},
		x0:      []float64{1.0},
		wantErr: false,
		want:    []float64{math.Sqrt(2)},
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
		x0:      []float64{1.0, 1.0},
		wantErr: false,
		want:    []float64{math.Sqrt(2), math.Sqrt(2)},
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
		x0:      []float64{0.5, 1.5, 2.5},
		wantErr: false,
		want:    []float64{1, 2, 3},
	},
	{
		name: "tri",
		equations: []func([]float64) float64{
			func(x []float64) float64 {
				return math.Sin(x[0]) - 0.5
			},
		},
		x0:      []float64{0.5},
		wantErr: false,
		want:    []float64{math.Pi / 6},
	},
	{
		name: "exp",
		equations: []func([]float64) float64{
			func(x []float64) float64 {
				return math.Exp(x[0]) - 2
			},
		},
		x0:      []float64{0.5},
		wantErr: false,
		want:    []float64{math.Log(2)},
	},
	{
		name:      "empty equations",
		equations: []func([]float64) float64{},
		x0:        []float64{1.0},
		wantErr:   true,
	},
	{
		name: "empty initial guess",
		equations: []func([]float64) float64{
			func(x []float64) float64 { return x[0]*x[0] - 2 },
		},
		x0:      []float64{},
		wantErr: true,
	},
}

// Тест решения системы нелинейных уравнений
func TestRoot(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Root(test.equations, test.x0)

			if err == nil == test.wantErr || err != nil != test.wantErr {
				t.Errorf("Root() error = %v, wantErr %v", err, test.wantErr)
				return
			}

			if !test.wantErr {
				if result == nil {
					t.Error("Root() returned nil result")
					return

				}
				for i := range result.X {
					if math.Abs(result.X[i]-test.want[i]) > 1e-6 {
						t.Errorf("Root().X[%v] = %f, want %f", i, result.X[i], test.want[i])
					}
				}
			}
		})
	}
}

// Бенчмарк решения системы нелинейных уравнений
func BenchmarkRoot(b *testing.B) {
	for _, test := range tests {
		if test.wantErr {
			continue
		}

		b.Run(test.name, func(b *testing.B) {
			// Сбрасываем таймер перед запуском бенчмарка
			b.ResetTimer()

			// Запускаем бенчмарк
			for i := 0; i < b.N; i++ {
				// Создаем копию начального приближения для каждого запуска
				x0Copy := make([]float64, len(test.x0))
				copy(x0Copy, test.x0)

				_, err := Root(test.equations, x0Copy)
				if err != nil {
					b.Errorf("Root() error = %v", err)
				}
			}
		})
	}
}
