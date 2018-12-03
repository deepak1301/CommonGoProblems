package main

import (
	"errors"
	"fmt"
)

func transpose(x [][]float32) [][]float32 {
	out := make([][]float32, len(x[0]))
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[0]); j++ {
			out[j] = append(out[j], x[i][j])
		}
	}
	return out
}

func dot(x, y [][]float32) ([][]float32, error) {
	if len(x[0]) != len(y) {
		return nil, errors.New("Can't do matrix multiplication")
	}

	out := make([][]float32, len(x), len(y[0]))
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(y[0]); j++ {
			for k := 0; k < len(x[0]); k++ {
				if len(out[i]) < 1 {
					out[i] = make([]float32, len(y))
				}
				out[i][j] += x[i][k] * y[k][j]
			}
		}
	}
	return out, nil
}

func main() {
	X := [][]float32{
		[]float32{1.0, 2.0, 3.0},
		[]float32{2.0, 2.0, 2.0},
		[]float32{-1.0, 0.0, 1.0},
	}

	out, _ := dot(X, transpose(X))
	fmt.Println(out)
}
