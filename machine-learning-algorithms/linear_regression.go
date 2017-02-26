package mlearn

import (
	"github.com/gonum/matrix/mat64"
)

//LeastSquares performs derives the "best fit" from the least squares method
func LeastSquares(X *mat64.Dense, y *mat64.Dense) mat64.Dense {
	var XT = X.T()

	var XTX mat64.Dense
	XTX.Mul(XT, X)

	var XTXInverse mat64.Dense
	_ = XTXInverse.Inverse(&XTX)

	var XTransposeY mat64.Dense
	XTransposeY.Mul(XT, y)

	var theta mat64.Dense
	theta.Mul(&XTXInverse, &XTransposeY)

	return theta
}

//CalculateTotalDelta takes the design matrix, the solution vector and the
//hypothesis and calculates the total delta
func CalculateTotalDelta(X *mat64.Dense, theta *mat64.Dense, y *mat64.Dense) float64 {
	rows, cols := X.Dims()

	delta := 0.0

	for i := 0; i < rows-1; i++ {
		trainingExample := mat64.NewDense(1, cols, X.RawRowView(i))

		var expectedValueMat mat64.Dense
		expectedValueMat.Mul(trainingExample, theta)
		expectedValue := expectedValueMat.RawRowView(0)[0]

		thisY := y.RawRowView(i)[0]

		thisDelta := (expectedValue - thisY) * (expectedValue - thisY)
		delta += thisDelta
	}

	return delta
}
