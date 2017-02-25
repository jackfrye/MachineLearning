package mlearn

import (
  "github.com/gonum/matrix/mat64"
)

func NormalEquation(X *mat64.Dense, y *mat64.Dense) mat64.Dense {
  var XT = X.T();

	var XTX mat64.Dense
	XTX.Mul(XT, X)

	var XTXInverse mat64.Dense
	_ = XTXInverse.Inverse(&XTX)

	var XTransposeY mat64.Dense
	XTransposeY.Mul(XT, y);

	var theta mat64.Dense
	theta.Mul(&XTXInverse, &XTransposeY);

  return theta
}
