package main

import(
	"bytes"
	"fmt"
	"io"
	"github.com/gonum/matrix/mat64"
	"github.com/jackfrye/MachineLearning/ReadFile"
)

func main() {
	read()
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

func read() {
	workBooks := ReadFile.ReadFile()

	fmt.Print(workBooks)
}
