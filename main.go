package main

import (
	"fmt"
	"path/filepath"
	"sort"

	"github.com/gonum/matrix/mat64"
	"github.com/jackfrye/MachineLearning/data"
	"github.com/jackfrye/MachineLearning/io"
	"github.com/jackfrye/MachineLearning/machine-learning-algorithms"
	"github.com/tonnerre/golang-pretty"
)

//main function to the project.
func main() {

	//Make absolute path filenames for your computer
	filenames := []string{
		makePath("res/ExportData.xlsx"),
		makePath("res/S&PStockPrices.xlsx"),
		makePath("res/US_GDP.xlsx")}

	//Read from the files and populate structure
	excelData := make([]map[string][]map[string]float64, 0, 10000)
	for _, filename := range filenames {
		excelData = append(excelData, read(filename))
	}

	//Flatten data to key by year
	historicalMarketData := data.JoinOn("Year", excelData)

	finalData := data.FilterIn(1977, 2015, historicalMarketData)

	// and then print with and without zero value elements.
	X, y := buildSystemOn("Price", finalData)

	XF := mat64.Formatted(X, mat64.Prefix("    "), mat64.Squeeze())
	fmt.Printf("\n\n%v", XF)

	yF := mat64.Formatted(y, mat64.Prefix("    "), mat64.Squeeze())
	fmt.Printf("\n\n%v", yF)

	//Now we perform linear regression using the least squares method
	// theta = ((X*X^T)^-1)(X^T*y)
	theta := mlearn.LeastSquares(X, y)
	delta := mlearn.CalculateTotalDelta(X, theta, y)

	fTheta := mat64.Formatted(theta, mat64.Prefix("    "), mat64.Squeeze())
	fmt.Printf("\n\n%v", fTheta)
	pretty.Print("\n\nDELTA", delta)

	theta2 := mlearn.GradientDescent(X, y, 800, .001)
	delta2 := mlearn.CalculateTotalDelta(X, theta2, y)

	fTheta2 := mat64.Formatted(theta2, mat64.Prefix("    "), mat64.Squeeze())
	fmt.Printf("\n\n%v\n\n", fTheta2)
	pretty.Print("\n\nDELTA2", delta2)

}

//test1 tests the mat64 library. Playground for features.
func test1() {
	var A = mat64.NewDense(2, 2, []float64{0.1, 3.4, 8.5, 22.0})

	var b = mat64.NewVector(2, []float64{7.15, 73.75})

	var x = mat64.NewVector(2, make([]float64, 2))

	x.SolveVec(A, b)
}

//read is a wrapper for reading files. Should be deprecated.
func read(filename string) map[string][]map[string]float64 {
	workBooks := io.ReadFile(filename)

	return workBooks
}

func makePath(filename string) string {
	a, err := filepath.Abs(filename)
	if err != nil {
		panic(err)
	}

	return a
}

//buildSystemOn takes data and builds a design matrix and solution.
func buildSystemOn(key string, data map[int]map[string]float64) (*mat64.Dense, *mat64.Dense) {
	rows := len(data)
	cols := 0

	length := rows

	designMatrixData := make([]float64, 0, 100000)
	solutionData := make([]float64, 0, 10000)

	for _, associatedData := range data {
		cols = len(associatedData)
		keys := make([]string, 0, len(associatedData))
		/*ensure you are going through the keys in the same order
		every tme */
		for k := range associatedData {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		designMatrixData = append(designMatrixData, 1)

		for _, sortedKey := range keys {
			if sortedKey != key {
				designMatrixData = append(designMatrixData, associatedData[sortedKey])
			} else {
				solutionData = append(solutionData, associatedData[key])
			}
		}
	}

	designMatrix := mat64.NewDense(rows, cols, designMatrixData)
	y := mat64.NewDense(length, 1, solutionData)

	return designMatrix, y
}
