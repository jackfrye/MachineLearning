//Package mlearn contains functions created for the purposes of machine learning
package mlearn

import (
	"github.com/gonum/matrix/mat64"
)

//LeastSquares finds the least squares solution of some design matrix X
//and some data points y.
func LeastSquares(X *mat64.Dense, y *mat64.Dense) *mat64.Dense {
	var XT = X.T()

	var XTX mat64.Dense
	XTX.Mul(XT, X)

	var XTXInverse mat64.Dense
	_ = XTXInverse.Inverse(&XTX)

	var XTransposeY mat64.Dense
	XTransposeY.Mul(XT, y)

	var theta mat64.Dense
	theta.Mul(&XTXInverse, &XTransposeY)

	return &theta
}

//GradientDescent will descend toward the best fit for the system X, producing
//the output theta
func GradientDescent(X *mat64.Dense, y *mat64.Dense, numIter int, alpha float64) *mat64.Dense {
	r, c := X.Dims()

	thetaPrep := make([]float64, c, c)
	for i := 0; i < c; i++ {
		thetaPrep[i] = 0.0
	}

	theta := mat64.NewDense(c, 1, thetaPrep)

	tempTheta := mat64.NewDense(c, 1, thetaPrep)

	for i := 0; i < numIter; i++ {
		for j := 0; j < c; j++ {
			sigma := 0.0
			for k := 0; k < r; k++ {
				var XRowTheta mat64.Dense
				XRowTheta.Mul(mat64.NewDense(1, c, X.RawRowView(k)), theta)

				XRowThetaLessY := XRowTheta.At(0, 0) - y.At(k, 0)

				total := XRowThetaLessY * X.At(k, j)

				sigma += total

			}
			newTheta := theta.At(j, 0) - (alpha)*sigma
			tempTheta.Set(j, 0, newTheta)

		}
		theta = tempTheta
	}

	return theta
}

//CalculateTotalDelta takes the design matrix, the solution vector and the
//hypothesis and calculates the total delta.
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
