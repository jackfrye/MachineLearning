package main

import (
	"path/filepath"
	"sort"

	"github.com/gonum/matrix/mat64"
	"github.com/jackfrye/MachineLearning/data"
	"github.com/jackfrye/MachineLearning/machine-learning-algorithms"
	"github.com/jackfrye/MachineLearning/readfile"
	"github.com/tonnerre/golang-pretty"
)

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

	X, y := buildSystemOn("Price", finalData)

	//Now we calculate the solution using the normal equation
	// theta = ((X*X^T)^-1)(X^T*y)
	theta := mlearn.NormalEquation(X, y)

	pretty.Print("HYPOTHESIS", theta)

	delta := mlearn.CalculateTotalDelta(X, &theta, y)

	pretty.Print("DELTA", delta)

}

func test1() {
	var A = mat64.NewDense(2, 2, []float64{0.1, 3.4, 8.5, 22.0})

	var b = mat64.NewVector(2, []float64{7.15, 73.75})

	var x = mat64.NewVector(2, make([]float64, 2))

	x.SolveVec(A, b)
}

func read(filename string) map[string][]map[string]float64 {
	workBooks := readfile.ReadFile(filename)

	return workBooks
}

func makePath(filename string) string {
	a, err := filepath.Abs(filename)
	if err != nil {
		panic(err)
	}

	return a
}

func buildSystemOn(key string, data map[int]map[string]float64) (*mat64.Dense, *mat64.Dense) {
	pretty.Print(data)

	rows := len(data)
	cols := 0

	length := rows

	designMatrixData := make([]float64, 0, 100000)
	solutionData := make([]float64, 0, 10000)

	for _, associatedData := range data {
		cols = len(associatedData) - 1
		keys := make([]string, 0, len(associatedData))
		/*ensure you are going through the keys in the same order
		every tme */
		for k := range associatedData {
			keys = append(keys, k)
		}
		sort.Strings(keys)

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
