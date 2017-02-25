package main

import(
	"bytes"
	"fmt"
	"io"
	"sort"
	"path/filepath"
	"github.com/gonum/matrix/mat64"
	"github.com/tonnerre/golang-pretty"
	"github.com/jackfrye/MachineLearning/readfile"
	"github.com/jackfrye/MachineLearning/data"
	"github.com/jackfrye/MachineLearning/machine-learning-algorithms"


)

func main() {

	//Make absolute path filenames for your computer
	filenames := []string{
		makePath("res/ExportData.xlsx"),
		makePath("res/S&PStockPrices.xlsx"),
		makePath("res/US_GDP.xlsx")};

		//Read from the files and populate structure
	excelData := make([]map[string][]map[string]float64, 0, 10000);
	for _, filename := range filenames {
		excelData = append(excelData, read(filename));
	}

	//Flatten data to key by year
	historicalMarketData := data.JoinOn("Year", excelData);

	finalData := data.FilterIn(1977, 2015, historicalMarketData);

	X, y := buildSystemOn("Price", finalData);

	//Now we calculate the solution using the normal equation
	// theta = ((X*X^T)^-1)(X^T*y)
	theta := mlearn.NormalEquation(X, y);

	pretty.Print("HYPOTHESIS", theta)

}

func test1() {
	var A *mat64.Dense = mat64.NewDense(2, 2, []float64{0.1, 3.4, 8.5, 22.0});

	var b *mat64.Vector = mat64.NewVector(2, []float64{7.15, 73.75});

	var x *mat64.Vector = mat64.NewVector(2, make([]float64, 2));

	x.SolveVec(A, b);

	byte_out, err := x.MarshalBinary();

	if(err == nil) {
		fmt.Println("no error");

		buf := bytes.NewBuffer(byte_out);

		w := io.MultiWriter(buf);

		fmt.Fprint(w);
	}
	//fmt.Printf("%+v\n", *x)
}

func read(filename string)  map[string][]map[string]float64 {
	workBooks := readfile.ReadFile(filename);

	return workBooks
}

func makePath(filename string) string {
	a, err := filepath.Abs(filename);
	if err == nil {
		return a
	} else {
		panic(err)
	}
}

func buildSystemOn(key string, data map[int]map[string]float64) (*mat64.Dense, *mat64.Dense) {
		rows := len(data);
		cols := 0;

		length := rows;

		designMatrixData := make([]float64, 0, 100000);
		solutionData := make([]float64, 0, 10000);

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
					designMatrixData = append(designMatrixData, associatedData[key])
				} else {
					solutionData = append(solutionData, associatedData[key])
				}
			}
		}

		designMatrix := mat64.NewDense(rows, cols, designMatrixData);
		y := mat64.NewDense(length, 1, solutionData);

		return designMatrix, y
}
