package main

import(
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"github.com/gonum/matrix/mat64"
	"github.com/tonnerre/golang-pretty"
	"github.com/jackfrye/MachineLearning/ReadFile"
	"github.com/jackfrye/MachineLearning/Data"

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
	historicalMarketData := Data.JoinOn("Year", excelData)

	finalData := Data.FilterIn(1977, 2015, historicalMarketData)

	pretty.Print(finalData)

	//test1()
}

func test1() {
	var A *mat64.Dense = mat64.NewDense(2, 2, []float64{0.1, 3.4, 8.5, 22.0})

	var b *mat64.Vector = mat64.NewVector(2, []float64{7.15, 73.75})

	var x *mat64.Vector = mat64.NewVector(2, make([]float64, 2))

	x.SolveVec(A, b)

	byte_out, err := x.MarshalBinary()

	if(err == nil) {
		fmt.Println("no error")

		buf := bytes.NewBuffer(byte_out)

		w := io.MultiWriter(buf)

		fmt.Fprint(w)
	}
	//fmt.Printf("%+v\n", *x)
}

func read(filename string)  map[string][]map[string]float64 {
	workBooks := ReadFile.ReadFile(filename)

	return workBooks
}

func makePath(filename string) string {
	a, err := filepath.Abs(filename)
	if err == nil {
		return a
	} else {
		panic(err)
	}
}
