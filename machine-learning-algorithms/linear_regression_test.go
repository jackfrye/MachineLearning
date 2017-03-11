package mlearn

import (
	"testing"

	"github.com/gonum/matrix/mat64"
)

func TestLeastSquares(t *testing.T) {
	a := mat64.NewDense(3, 3, []float64{1.0, 2.0, 3.0, 2.0, 5.0, 9.0, 3.0, 5.0, 10.0})
	b := mat64.NewDense(3, 1, []float64{6.0, 16.0, 18.0})
	theta := LeastSquares(a, b)
	expectedOutput := []float64{1.0, 1.0, 1.0}

	max := 0.0
	r, _ := theta.Dims()
	for i := 0; i < r; i++ {
		if theta.At(i, 0) > float64(max) {
			max = theta.At(i, 0)
		}
	}
	for i := 0; i < r; i++ {
		theta.Set(i, 0, theta.At(i, 0)/max)
	}

	for i, value := range expectedOutput {
		if value-theta.At(i, 0) > .1 {
			t.Errorf("delta to high for index index=%d delta=%f", i, value-theta.At(i, 0))
		}
	}

}

func TestGradientDescent(t *testing.T) {
	a := mat64.NewDense(3, 3, []float64{1.0, 2.0, 3.0, 2.0, 5.0, 9.0, 3.0, 5.0, 10.0})
	b := mat64.NewDense(3, 1, []float64{6.0, 16.0, 18.0})
	theta := GradientDescent(a, b, 400, .01)
	expectedOutput := []float64{1.0, 1.0, 1.0}
	for i, value := range expectedOutput {
		if value-theta.At(i, 0) > .1 {
			t.Errorf("delta to high for index index=%d delta=%f", i, value-theta.At(i, 0))
		}
	}
}
